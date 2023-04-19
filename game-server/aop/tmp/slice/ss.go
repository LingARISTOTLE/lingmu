package slice

import "fmt"

func S() {
	var s = []int{1, 2, 3, 5, 8}
	fmt.Println(cap(s))
	//前面两个是获取切片的范围，左闭右开，右边的是切片的最大长度
	newne := s[1:2:4]
	fmt.Println(newne)
	fmt.Println(cap(newne))

}
