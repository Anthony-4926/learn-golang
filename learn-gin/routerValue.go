package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/gomodule/redigo/redis"

func RouterValue() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	//	get请求传值
	r.GET("/getValue", func(c *gin.Context) {
		if rec, ok := redis.Dial("tcp", ":6379"); ok == nil {
			defer rec.Close()
			key := c.Query("key")
			v, _ := rec.Do("GET", key)
			tt := fmt.Sprintf("%s", v)
			fmt.Printf("%s\n", v)
			c.JSON(http.StatusOK, gin.H{
				key: tt,
			})
		} else {
			c.String(http.StatusInternalServerError, "请求失败")
		}
	})

	//	post请求传值
	r.GET("/postValue", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("/addKey", func(c *gin.Context) {
		if rec, ok := redis.Dial("tcp", ":6379"); ok == nil {
			defer rec.Close()
			value := c.Query("name")
			v, _ := rec.Do("SET", "name", value)
			c.String(http.StatusOK, fmt.Sprintf("%s", v))
		} else {
			c.String(http.StatusInternalServerError, "请求失败")
		}
	})

	r.Run()

}
