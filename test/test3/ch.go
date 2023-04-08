package main

import (
	"fmt"
	"sync"
)

func main() {
	w2.Add(2)
	ChDemo()
	ChDemo2()
	w2.Wait()
	demo3()

}

func demo3() {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c1 <- 3
	go demo3Sub(c1, c2)
	fmt.Println(<-c2)
	close(c1)
	close(c2)
}

func demo3Sub(in chan int, out chan int) {
	for i := range in {
		out <- i + 1
	}
}

var ch2 = make(chan int)    //无缓冲
var ch3 = make(chan int, 5) //有缓冲
var w sync.WaitGroup
var w2 sync.WaitGroup

func ChDemo() {
	w.Add(1)
	c := ChDemoSub()
	go func() {
		w.Wait()
		fmt.Println(<-c)
		close(c)
		w2.Done()
	}()
}

func ChDemoSub() chan int {
	var ch1 chan int
	ch1 = make(chan int, 2) //you缓冲
	ch1 <- 9
	w.Done()
	return ch1
}

func ChDemo2() {
	ch3 <- 7
	go ChDemo2Sub()
}

func ChDemo2Sub() {
	val := <-ch3
	fmt.Println(val)
	close(ch3)
	w2.Done()
}
