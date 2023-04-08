package main

import "fmt"

type Dog struct {
	Nick string `json:"nick"`
	Age  int    `json:"age"`
}

func (d *Dog) Eat(i interface{}) {
	//TODO implement me
	fmt.Println(d.Nick + "eat")
	//panic("implement me")
}

func (d *Dog) Name() string {
	//TODO implement me
	//fmt.Println(d.Nick)
	//panic("implement me")

	return d.Nick
}
