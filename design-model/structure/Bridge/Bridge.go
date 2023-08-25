package Bridge

import "fmt"

type Phone interface {
	call()
}

type Xiaomi struct {
	color Color
	pixel Pixel
}

func (x Xiaomi) call() {
	fmt.Println("小米手机打电话")
}

func NewXiaomi(color Color, pixel Pixel) *Xiaomi {
	return &Xiaomi{color: color, pixel: pixel}
}

type Huawei struct {
	color Color
	pixel Pixel
}

func (x Huawei) call() {
	fmt.Println("华为手机打电话")
}

func NewHuawei(color Color, pixel Pixel) *Huawei {
	return &Huawei{color: color, pixel: pixel}
}

type Color interface {
	showColor()
}
type Pixel interface {
	showPixel()
}

type Black struct{}

func (c Black) showColor() {
	fmt.Println("黑色")
}

type Gold struct{}

func (c Gold) showColor() {
	fmt.Println("金色")
}

type P800 struct{}

func (p P800) showPixel() {
	fmt.Println("800万像素")
}

type P1000 struct{}

func (p P1000) showPixel() {
	fmt.Println("1000万像素")
}
