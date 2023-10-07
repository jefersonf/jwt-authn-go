package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jefersonf/jwt-authn-go/env"
	"github.com/jefersonf/jwt-authn-go/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// get the email/pass off request body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	// create the user
	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}

	conn := env.Conn()
	dbHandler := conn()

	result := dbHandler.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"message": "signing up completed",
	})
}
