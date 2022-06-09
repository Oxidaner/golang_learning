package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})
	r.POST("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello Post")
	})
	r.PUT("/Myput", func(context *gin.Context) {
		context.String(http.StatusOK, "hello Put")
	})
	//监听端口默认为8080
	r.Run(":8000")
}
