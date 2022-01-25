package app

import "github.com/gin-gonic/gin"

func RunServer() {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	r.Run()
}
