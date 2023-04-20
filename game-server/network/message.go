package network

import "sync"

/*
Message
@Description: 网络包
*/
type Message struct {
	Id   uint64 // message序号
	Data []byte //message数据
}

var msgPool = sync.Pool{
	New: func() interface{} {
		return &Message{}
	},
}

// GetPooledMessage gets a pooled Message.
func GetPooledMessage() *Message {
	return msgPool.Get().(*Message)
}

// FreeMessage puts a Message into the pool.
func FreeMessage(msg *Message) {
	if msg != nil && len(msg.Data) > 0 {
		ResetMessage(msg)
		msgPool.Put(msg)
	}
}

// ResetMessage reset a Message
func ResetMessage(m *Message) {
	m.Data = m.Data[:0]
}
