package main

import "fmt"

/*
测试泛型
*/

func main() {
	t := AddT[float32](1.0, 2.2)
	fmt.Println(t)
}

func Add(a int, b int) int {
	return a + b
}

func AddFloat32(a float32, b float32) float32 {
	return a + b
}

func AddString(a string, b string) string {
	return a + b
}

func AddT[T int | int32 | float32 | string | float64](a T, b T) T {
	return a + b
}
