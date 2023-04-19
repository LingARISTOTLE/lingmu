package lock

import (
	"fmt"
	"sync"
	"testing"
)

type TBZ struct {
	L       Lock
	Counter int64
}

func (t *TBZ) DoBz() {

}

func TestTryLock(t *testing.T) {
	tt := TBZ{}
	tt.L = NewLock()
	w := sync.WaitGroup{}
	//mutex := sync.Mutex{}
	for tt.Counter < 10000 {
		w.Add(1)
		go func() {
			defer w.Done()
			if !tt.L.Lock() {
				return
			}
			//mutex.Lock()

			tt.Counter++
			//mutex.Unlock()
			tt.L.Unlock()
		}()
	}
	w.Wait()
	fmt.Println(tt.Counter)
	//=== RUN   TestTryLock
	//10002
	//--- PASS: TestTryLock (0.01s)
	//达到锁的功能即可
	//超过10000的原因是tt.Counter < 10000并非原子性
}
