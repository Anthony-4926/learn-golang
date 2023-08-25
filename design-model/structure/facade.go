package main

import "fmt"

type moYu struct{}

func (y moYu) open() {
	fmt.Println("打开摸鱼软件")
}

type IDE struct{}

func (i IDE) open() {
	fmt.Println("打开IDE")
}

type browser struct{}

func (b browser) open() {
	fmt.Println("打开浏览器")
}

type terminal struct{}

func (t terminal) open() {
	fmt.Println("打开终端控制台窗口")
}

type facade struct {
	yu         moYu
	ide        IDE
	b          browser
	t1, t2, t3 terminal
}

func (f facade) open() {
	f.yu.open()
	f.ide.open()
	f.b.open()
	f.t1.open()
	f.t2.open()
	f.t3.open()
}

func NewFacade() *facade {
	return &facade{
		yu:  moYu{},
		ide: IDE{},
		b:   browser{},
		t1:  terminal{},
		t2:  terminal{},
		t3:  terminal{},
	}
}
