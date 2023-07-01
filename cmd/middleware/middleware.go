package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("token noy found in environment variable")
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, errors.New("empty token"))
			c.Abort()
			return
		}

		if token != requiredToken {
			c.JSON(http.StatusUnauthorized, errors.New("invalid token"))
			c.Abort()
			return
		}

		c.Next()
	}
}
