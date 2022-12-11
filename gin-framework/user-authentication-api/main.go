package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var user = []User{
	{Username: "vn", Password: "123"},
	{Username: "user", Password: "password"},
}

func Create_user(c *gin.Context) {
	var profile User
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user = append(user, profile)
	c.JSON(http.StatusOK, gin.H{"status": "user created"})
}

func login(c *gin.Context) {
	var profile User
	if err := c.BindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, u := range user {
		fmt.Println(u.Username, u.Password)
		if profile.Username == u.Username || profile.Password == u.Password {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			return
		}
	}
	c.IndentedJSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
}

func main() {
	r := gin.Default()
	r.POST("/login", login)
	r.POST("/create_user", Create_user)
	r.Run()
}
