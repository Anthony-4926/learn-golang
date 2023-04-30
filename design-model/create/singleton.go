package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Singleton 单例类
type Singleton struct{}

// 全局单例对象
var instance *Singleton

var sema int64

var once sync.Once

// GetInstance 对外提供获取单例对象的方法
func GetInstance() *Singleton {
	if atomic.CompareAndSwapInt64(&sema, 0, 1) {
		fmt.Println("---")
		instance = &Singleton{}
	}

	return instance
}
