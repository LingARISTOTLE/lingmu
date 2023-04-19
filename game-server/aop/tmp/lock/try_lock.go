package lock

/*
Lock
@Description: 基于管道实现的锁
*/
type Lock struct {
	c chan struct{}
}

// NewLock generate a try lock
func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

/*
Lock
@Description: 多协程同时读取管道的值，能读取到的获得锁
@receiver l
@return bool
*/
func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

/*
Unlock
@Description: 释放锁，当管道中有值了就释放锁
@receiver l
*/
func (l Lock) Unlock() {
	l.c <- struct{}{}
}
