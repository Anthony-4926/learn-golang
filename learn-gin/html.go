package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func StatCost(c *gin.Context) {
	start := time.Now()
	// 调用该请求的剩余处理程序
	c.Next()
	// 计算耗时
	cost := time.Since(start)
	fmt.Println(cost)
}

func main() {
	r := gin.New()

	shopGroup := r.Group("/shop")
	shopGroup.Use(StatCost)
	{
		shopGroup.GET("/index", func(c *gin.Context) {})
	}

	r.Run()
}
