package main

type Animal interface {
	Eat(interface{}) //定义接口方法
	Name() string
}

type Base struct {
}

func (b *Base) Eat(i interface{}) {
	//TODO implement me
	panic("implement me")
}

func (b *Base) Name() string {
	//TODO implement me
	panic("implement me")
}
