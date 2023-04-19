package network

import (
	"fmt"
	spoor "github.com/LingARISTOTLE/lingolog"
	"net"
	"os"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

type Server struct {
	pid            int64                    //进程号
	Addr           string                   //网络地址
	MaxConnNum     int                      //最大连接数
	listen         *net.TCPListener         //TCP监听端口
	connSet        map[net.Conn]interface{} //连接集合
	counter        int64                    //连接数量
	idCounter      int64                    //id号
	mutexConn      sync.Mutex               //锁
	wgLn           sync.WaitGroup           //
	wgConn         sync.WaitGroup           //
	connBuffSize   int                      //连接缓冲大小
	logger         *spoor.Spoor             //日志
	MessageHandler func(packet *Packet)     //消息处理
}

/*
NewServer
@Description: 创建服务对象
@param address
@param network
@return *Server
*/
func NewServer(addr string, maxConnNum int, buffSize int, logger *spoor.Spoor) *Server {
	s := &Server{
		Addr:         addr,
		MaxConnNum:   maxConnNum,
		connBuffSize: buffSize,
		logger:       logger,
	}
	s.Init()
	return s
}

/*
Init
@Description: 初始化服务器
@receiver s
*/
func (s *Server) Init() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", s.Addr)
	if err != nil {
		fmt.Printf("[net] addr resolve error", tcpAddr, err)
	}

	listen, err := net.ListenTCP("tcp6", tcpAddr)
	if err != nil {
		fmt.Printf("%v", err)
	}

	if s.MaxConnNum <= 0 {
		s.MaxConnNum = 100
		fmt.Printf("invalid MaxConnNum, reset to %v", s.MaxConnNum)
	}
	s.listen = listen
	s.connSet = make(map[net.Conn]interface{})
	s.counter = 1
	s.idCounter = 1
	s.pid = int64(os.Getpid())
	fmt.Printf("Server Listen %s", s.listen.Addr().String())
}

/*
Run
@Description: 启动服务器
@receiver s
*/
func (s *Server) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("[net] panic", err, "\n", string(debug.Stack()))
		}
	}()

	s.wgLn.Add(1)

	defer s.wgLn.Done()

	var tempDelay time.Duration

	for {
		conn, err := s.listen.AcceptTCP()
		//设置超时时间
		if err != nil {
			if _, ok := err.(net.Error); ok {
				if tempDelay == 0 {
					tempDelay = 5 * time.Microsecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}

				fmt.Printf("accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return
		}

		tempDelay = 0

		//判断连接数
		if atomic.LoadInt64(&s.counter) >= int64(s.MaxConnNum) {
			//达到最大连接数，那么关闭
			conn.Close()
			fmt.Printf("too many connections %v", atomic.LoadInt64(&s.counter))
			continue
		}

		//与该用户生成连接
		tcpConnX, err := NewTcpConnX(conn, s.connBuffSize, s.logger)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		//注册连接
		s.addConn(conn, tcpConnX)
		tcpConnX.Impl = s

		//启动读写协程
		s.wgConn.Add(1)
		go func() {
			tcpConnX.Connect()
			s.removeConn(conn, tcpConnX)
			s.wgConn.Done()
		}()
	}

}

/*
Close
@Description: 关闭监听以及连接，关闭过程原子操作
@receiver s
*/
func (s *Server) Close() {
	s.listen.Close()
	s.wgLn.Wait()

	s.mutexConn.Lock()
	for conn := range s.connSet {
		conn.Close()
	}
	s.connSet = nil
	s.mutexConn.Unlock()
	s.wgConn.Wait()
}

/*
addConn
@Description: 将连接添加到server的集合，统一管理
@receiver s
@param conn
@param tcpConnX
*/
func (s *Server) addConn(conn net.Conn, tcpConnX *TcpConnX) {
	s.mutexConn.Lock()
	//连接数++
	atomic.AddInt64(&s.counter, 1)
	s.connSet[conn] = conn
	nowTime := time.Now().Unix()
	//当前连接数++
	idCounter := atomic.AddInt64(&s.idCounter, 1)
	//计算连接id
	connId := (nowTime << 32) | (s.pid << 24) | idCounter
	tcpConnX.ConnID = connId
	s.mutexConn.Unlock()
	//调用连接时处理方法
	tcpConnX.OnConnect()
}

/*
removeConn
@Description: 移除管理的连接
@receiver s
@param conn
@param tcpConn
*/
func (s *Server) removeConn(conn net.Conn, tcpConn *TcpConnX) {
	tcpConn.Close()
	s.mutexConn.Lock()
	atomic.AddInt64(&s.counter, -1)
	delete(s.connSet, conn)
	s.mutexConn.Unlock()
}

func (s *Server) OnMessage(message *Message, conn *TcpConnX) {
	s.MessageHandler(&Packet{
		Msg:  message,
		Conn: conn,
	})
}

func (s *Server) OnClose() {

}

func (s *Server) OnConnect() {

}

//获取tcp连接地址
//resolveTCPAddr, err := net.ResolveTCPAddr("tcp6", address)
//if err != nil {
//	panic(err)
//}
//
////获取连接监听器
//tcpListener, err := net.ListenTCP("tcp6", resolveTCPAddr)
//if err != nil {
//	panic(err)
//}
//
//s := &Server{}
//s.listener = tcpListener

//listener        net.Listener         //服务器请求监听器
//OnSessionPacket func(*SessionPacket) //处理网络包
//Address         string               //服务器监听端口
//network  string

//
//resolveTCPAddr, err := net.ResolveTCPAddr("tcp6", s.Address)
//if err != nil {
//	panic(err)
//}
//
//tcpListener, err := net.ListenTCP("tcp6", resolveTCPAddr)
//if err != nil {
//	panic(err)
//}
//
//s.listener = tcpListener
//
//for {
//	//循环监听
//	conn, err := s.listener.Accept()
//	fmt.Println("获取连接")
//	if err != nil {
//		if _, ok := err.(net.Error); ok {
//			fmt.Println(err)
//			continue
//		}
//	}
//
//	//生成session
//	newSession := NewSession(conn)
//	SessionMgrInstance.AddSession(newSession)
//	//启动用户会话，由于内部使用协程启动了读写，所以没必要go出去
//	newSession.Run()
//	SessionMgrInstance.DelSession(newSession.UId)
//}
