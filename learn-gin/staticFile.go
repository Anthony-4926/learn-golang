package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StaticFile() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.Static("/static", "./static")
	r.GET("/static", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":  "学习gin框架",
			"imgSrc": "/static/images/img.png",
		})
	})
	r.Run()
}
