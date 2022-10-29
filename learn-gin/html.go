package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Html() {
	r := gin.Default()
	r.Static("./img", "./static")
	r.LoadHTMLGlob("./templates/*")
	r.GET("/", func(c *gin.Context) {
		stru := &student{Name: "zhaoxin", Age: 23}
		list := []string{"list1", "list2", "list3", "list4"}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"transfer": "这是我从后台传递过来的数据",
			"stru":     stru,
			"list":     list,
		})
	})
	r.Run()
}
