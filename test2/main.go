package main

import "fmt"

func main() {
	d := &Dog{
		Nick: "WangCai",
		Age:  0,
	}

	fmt.Println(d)
	name := d.Name()
	fmt.Println("---------------------")
	fmt.Println(name)

}
