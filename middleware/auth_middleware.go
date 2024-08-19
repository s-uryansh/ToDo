package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("session_id")
		if err != nil {
			log.Println("error getting token from cookie: ", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}
		// Verify the token and set it in the context
		c.Set("token", token)
		c.Next()
	}
}
