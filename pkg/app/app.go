package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "HEAD", "OPTIONS", "POST", "PUT"},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		AllowOrigins:     []string{"http://localhost:3000"},
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(corsConfig))

	r.GET("/hello", helloHandler)
	r.POST("/login", loginHandler)
	r.GET("/users", listUserHandler)
	r.GET("/user/:email", userDetails)

	r.Run()
}
