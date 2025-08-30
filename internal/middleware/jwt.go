package middleware

import (
	//"auth-service/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireNoAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        token, _ := c.Cookie("token")
        if token != "" {
            // Если уже залогинен, редирект на главную
            c.Redirect(http.StatusFound, "/catalogue")
            c.Abort()
            return
        }
        c.Next()
    }
}