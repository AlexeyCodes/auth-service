package handlers

import (
	"net/http"

	"auth-service/internal/config"
	"auth-service/internal/models"
	"auth-service/internal/repos"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Name or Password"})
		return
	}

	if user.Role == "" {
		user.Role = "user"
	}

	if err := repos.CreateUserInDB(config.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = ""

	// После успешной регистрации редирект на login
	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func LoginHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Name or Password"})
		return
	}

	token, err := repos.LoginUser(config.DB, &user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Сохраняем токен в HttpOnly cookie
	c.SetCookie("token", token, 3600*24*30, "/", "", false, true) // 30 дней
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func ProfileHandler(c *gin.Context) {
	userID := c.GetString("user_id")
	role := c.GetString("role")

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"role":    role,
	})
}

// Группируем роуты
func AuthRoutes(rt *gin.RouterGroup) {
	rt.POST("/login", LoginHandler)
	rt.POST("/register", RegisterHandler)
}
