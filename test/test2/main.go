package main

import "fmt"

func main() {
	//d := &Dog{
	//	Nick: "WangCai",
	//	Age:  0,
	//}
	//
	//fmt.Println(d)
	//name := d.Name()
	//fmt.Println("---------------------")
	//fmt.Println(name)

	//s := S{}
	//f(&s) //4

	//不难看出 interface 的变量中存储的是实现了 interface 的类型的对象值，这种能力是 duck typing。
	/*
		在使用 interface 时不需要显式在 struct 上声明要实现哪个 interface ，
		只需要实现对应 interface 中的方法即可，go 会自动进行 interface 的检查，
		并在运行时执行从其他类型到 interface 的自动转换，即使实现了多个 interface，
		go 也会在使用对应 interface 时实现自动转换，这就是 interface 的魔力所在。
	*/
	s := S{}
	var i I //声明 i
	i = &s  //赋值 s 到 i
	fmt.Println(i.Get())

}

// 1
type I interface {
	Get() int
	Set(int)
}

// 2
type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s *S) Set(age int) {
	s.Age = age
}

// 3
func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

//func main() {
//	s := S{}
//	f(&s)  //4
//}
