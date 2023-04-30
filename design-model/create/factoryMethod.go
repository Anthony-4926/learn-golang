package main

import "fmt"

type Huawei struct {
	producer string
	brand    string
}

func (p Huawei) Call() {
	fmt.Println(p.producer, "代工的", p.brand, "正在打电话")
}

type Xiaomi struct {
	producer string
	brand    string
}

func (p Xiaomi) Call() {
	fmt.Println(p.producer, "代工的", p.brand, "正在打电话")
}

// Factory 工厂接口
type Factory interface {
	CreateXiaomi() Phone
	CreateHuawei() Phone
}

// Biyadi 具体代工工厂
type Biyadi struct{}

func (f Biyadi) CreateXiaomi() Phone {
	return &Xiaomi{"比亚迪", "小米"}
}

func (f Biyadi) CreateHuawei() Phone {
	return &Huawei{"比亚迪", "华为"}
}

// Fushikang 具体代工工厂
type Fushikang struct{}

func (f Fushikang) CreateXiaomi() Phone {
	return &Xiaomi{"富士康", "小米"}
}

func (f Fushikang) CreateHuawei() Phone {
	return &Huawei{"富士康", "华为"}
}

//func main() {
//	var factory Factory
//	factory = new(Biyadi)
//	xiaomi := factory.CreateXiaomi()
//	huawei := factory.CreateHuawei()
//	xiaomi.Call()
//	huawei.Call()
//
//	factory = new(Fushikang)
//	xiaomi = factory.CreateXiaomi()
//	huawei = factory.CreateHuawei()
//	xiaomi.Call()
//	huawei.Call()
//}
