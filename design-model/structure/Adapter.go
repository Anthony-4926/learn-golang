package main

import "fmt"

type V5 interface {
	use5V()
}

type Phone struct{}

func (p Phone) charge(v5 V5) {
	v5.use5V()
}

type PowerSupply struct {
}

func (p PowerSupply) use220V() {
	fmt.Println("正在用220V电压")
}

type Adapter struct {
	PowerSupply
}

func (a Adapter) use5V() {
	fmt.Println("适配器将调用220V电压")
	a.PowerSupply.use220V()
	fmt.Println("适配器对220V电压进行转换")
	fmt.Println("正在用5V电压")
}
