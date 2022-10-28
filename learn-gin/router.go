package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"message": "这是我的第一个gin代码",
		})
	})
	r.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, student{Name: "zhaoxin", Age: 24})
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "我是post请求的返回")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "我是delete请求的返回"})
	})
	r.PUT("/put", func(c *gin.Context) {
		c.String(http.StatusOK, "我是put请求的返回")
	})

	r.Run()
}
