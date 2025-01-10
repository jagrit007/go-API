package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-tasks-app-practice/internal/config"
	"go-tasks-app-practice/internal/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

// var JWTSecret = []byte("secret")

var JWTSecret = os.Getenv("JWT_SECRET")

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.FindUserByEmail(config.DB, input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials" + err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 6).Unix(), // expires in a generous 6 hours
	})

	tokenString, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, try again"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "message": "Token generated successfully"})

}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := models.FindUserByEmail(config.DB, input.Email); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email is already registered"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//user := map[string]string{
	//	"email":    input.Email,
	//	"password": string(hashedPassword),
	//}

	user := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := user.Create(config.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}

	//models.AddUser(user)

	c.JSON(http.StatusCreated, gin.H{
		"user":    user,
		"message": "User created",
	})

}
