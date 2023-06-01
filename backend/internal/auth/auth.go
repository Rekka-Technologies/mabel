package internal

import (
	"net/http"

	"github.com/Rekka-Technologies/mabel/backend/internal/models"
	"github.com/gin-gonic/gin"
)

type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register creates a new user in the database.
func Register(c *gin.Context) {
	var input AuthInput

	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user already exists
	exists, _ := models.CheckUserExists(input.Username)
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
		return
	}

	u := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	// Create the user record
	if err := u.Create(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}

// Login checks if the username and password are correct then returns
// a jwt token which is used for all authenticated requests.
func Login(c *gin.Context) {
	var input AuthInput

	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate the user and generate a jwt token
	token, err := models.LoginCheck(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
