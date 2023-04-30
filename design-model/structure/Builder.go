package main

import "fmt"

type component interface {
	open()
	close()
}

type file struct {
	name string
}

func (f file) open() {
	fmt.Printf("打开文件: %s\n", f.name)
}
func (f file) close() {
	fmt.Printf("关闭文件: %s\n", f.name)
}

type folder struct {
	name string
	// 文件夹中包含组件
	components []component
}

func (f folder) open() {
	fmt.Printf("打开文件夹: %s\n", f.name)
}
func (f folder) close() {
	fmt.Printf("关闭文件夹: %s\n", f.name)
}
