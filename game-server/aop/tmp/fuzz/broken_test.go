package fuzz

import (
	"fmt"
	"testing"
)

/*
FuzzBrokenMethod
@Description: 是否为fuzz方法
@param f
*/
func FuzzBrokenMethod(f *testing.F) {
	f.Fuzz(func(t *testing.T, Data string) {
		BrokenMethod(Data)
	})
}

func FuzzMod(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b int) {
		fmt.Println(a / b)
	})
}

func FuzzReverse(f *testing.F) {
	f.Fuzz(func(t *testing.T, a string) {
		Reverse(a)
	})
}
