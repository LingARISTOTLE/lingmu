package main

import "fmt"

func main() {
	SliceDemo()
}

func SliceDemo() {
	var s1 []int
	s2 := make([]int, 9) // 为s2分配长度为9的连续int空间，9是最大长度
	fmt.Println(s2)

	s1 = make([]int, 0, 9)
	println(s1)

	s1 = append(s1, 11)
	s1 = append(s1, 22)

	s1[0] = 1
	s1[1] = 2
	fmt.Println(s1[0:]) //打印从0到结尾的切片

	/**
	go数组有固定长度，当前分配了9实际上是允许使用到9，最大长度为9
	当使用append时如果当前的最大空间能够承受本次添加，那么就不扩容，数组允许长度增大（也就是我们声明的长度）
	当前的最大长度不足以给被允许长度扩容，那么最大长度扩容，默认为+1/4
	*/
	s1 = append(s1, 9) //将1追加到s1切片后面
	fmt.Println(s1)

	println("---", s1[0])

	fmt.Println("--------------")

	var m1 map[int]bool
	m1 = make(map[int]bool) // 创建map对象
	m1[1] = true
	m1[3] = true
	m1[5] = true
	m1[6] = true

	for i, b := range m1 {
		fmt.Println(i, b)
	}

	fmt.Println("--------------")

	for i, i2 := range s1 {
		fmt.Println(i, i2)
	}
}
