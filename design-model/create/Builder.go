package main

import "fmt"

type Phone interface {
	Call()
}

type Xiaomi11 struct{}

func (p Xiaomi11) Call() {
	fmt.Println("xiaomi11 打电话")
}

type Xiaomi12 struct{}

func (p Xiaomi12) Call() {
	fmt.Println("xiaomi112 打电话")
}

// IBuilder Builder接口
type IBuilder interface {
	cut()
	welde()
	installCamera()
	pushCover()
	getPhone() Phone
}

// BuilderMi11 小米11的Builder实现
type BuilderMi11 struct{}

func (b *BuilderMi11) cut() {
	fmt.Println("小米11，切割边框")
}

func (b *BuilderMi11) welde() {
	fmt.Println("小米11，焊接主板")
}

func (b *BuilderMi11) installCamera() {
	fmt.Println("小米11， 安装摄像头模组")
}

func (b *BuilderMi11) pushCover() {
	fmt.Println("小米11，压后盖板")
}
func (b *BuilderMi11) getPhone() Phone {
	return Xiaomi11{}
}

// BuilderMi12 小米12的Builder实现
type BuilderMi12 struct{}

func (b *BuilderMi12) cut() {
	fmt.Println("小米12，切割边框")
}

func (b *BuilderMi12) welde() {
	fmt.Println("小米12，焊接主板")
}

func (b *BuilderMi12) installCamera() {
	fmt.Println("小米12， 安装摄像头模组")
}

func (b *BuilderMi12) pushCover() {
	fmt.Println("小米12，压后盖板")
}
func (b *BuilderMi12) getPhone() Phone {
	return Xiaomi12{}
}

type Direcotr struct {
	builder IBuilder
}

func (d *Direcotr) Construct() Phone {
	d.builder.cut()
	d.builder.welde()
	d.builder.installCamera()
	d.builder.pushCover()
	return d.builder.getPhone()
}

func main() {
	builderMi11 := &BuilderMi11{}
	director11 := &Direcotr{builderMi11}
	mi11 := director11.Construct()
	mi11.Call()

	fmt.Println("-------------------------------")

	builderMi12 := &BuilderMi12{}
	director12 := &Direcotr{builderMi12}
	mi12 := director12.Construct()
	mi12.Call()
}
