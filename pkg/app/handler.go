package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello"})
}

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func loginHandler(c *gin.Context) {
	// fmt.Println("helllloooo1")

	var loginDetails login

	// body, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(body))
	// fmt.Println("helllloooo")

	if err := c.BindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad data"})
	}
	fmt.Println("hello")
	// fmt.Println(loginDetails.Email)
	if loginDetails.Email == "ashish" && loginDetails.Password == "dental" {
		c.JSON(200, gin.H{"message": "Login Success"})
	}
}
