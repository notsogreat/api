package app

import (
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

type user struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func loginHandler(c *gin.Context) {
	// fmt.Println("helllloooo1")

	var loginDetails login

	if err := c.BindJSON(&loginDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad data"})
	}
	if loginDetails.Email == "ashish" && loginDetails.Password == "dental" {
		c.JSON(200, gin.H{"message": "Login Success"})
	}
}

func listUserHandler(c *gin.Context) {
	// fmt.Println("helllloooo1")
	var userDetails = []user{
		{
			FirstName: "Ashish",
			LastName:  "Heda",
		},
		{
			FirstName: "Abhishek",
			LastName:  "Heda",
		},
	}

	c.JSON(200, userDetails)
}
