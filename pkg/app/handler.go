package app

import (
	"github.com/gin-gonic/gin"
)

// func helloHandler() (c *gin.Context) {
// 	return c.JSON(200)

// }

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}
