package network

import "sync"

type SessionMgr struct {
	Sessions map[int64]*Session
	Counter  int64
	Mutex    sync.Mutex
	Pid      int64
}

var (
	SessionMgrInstance SessionMgr //群居变量，一个项目只有一个
	onceInitSessionMgr sync.Once  //sync.Once赋予的Do回调只可以调用一次
)

/*
init
@Description: 初始化Session管理器，初始化全局变量SessionMgr，也就是将SessionMgr创建对象全局可用
*/
func init() {
	//Do方法使得该对象只会被初始化一次
	onceInitSessionMgr.Do(func() {
		SessionMgrInstance = SessionMgr{
			Sessions: make(map[int64]*Session),
			Counter:  0,
			Mutex:    sync.Mutex{},
		}
	})
}

/*
AddSession
@Description: 添加Session
@receiver sm
@param s
*/
func (sm *SessionMgr) AddSession(s *Session) {
	sm.Mutex.Lock()
	defer sm.Mutex.Unlock()
	if val := sm.Sessions[s.UId]; val != nil {
		if val.IsClose {
			sm.Sessions[s.UId] = s
		} else {
			return
		}
	}
}

/*
DelSession
@Description: 删除Session
@receiver sm
@param UId
*/
func (sm *SessionMgr) DelSession(UId int64) {
	delete(sm.Sessions, UId)
}
