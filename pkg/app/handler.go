package app

import (
	"net/http"
	"strings"

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
	Email     string `json:"email"`
	LastPaid  string `json:"lastpaid"`
	Remaining string `json:"remaining"`
	TotalPaid string `json:"totalpaid"`
	Service   string `json:"service"`
}

type details struct {
	Email string
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

func userDetails(c *gin.Context) {

	userEmail := strings.TrimSpace(c.Param("email"))

	var userData = make(map[string]user)
	userData["ashishheda766@gmail.com"] = user{
		FirstName: "Ashish",
		LastName:  "Heda",
		Email:     "ashishheda766@gmail.com",
		LastPaid:  "200",
		Remaining: "300",
		TotalPaid: "500",
		Service:   "Cleaning",
	}

	if userEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "no email provided"})
		return
	}

	if output, ok := userData[userEmail]; ok {
		c.JSON(200, output)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
