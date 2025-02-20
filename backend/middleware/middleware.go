package middleware

import (
	"strings"
	"student-api/auth"

	"github.com/gin-gonic/gin"
)

func AuthWithExceptions(excludedPaths []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestPath := c.Request.URL.Path

		// Check if the request path matches any of the excluded paths
		for _, excludedPath := range excludedPaths {
			if strings.HasPrefix(requestPath, excludedPath) {
				c.Next()
				return
			}
		}

		tokenString := c.GetHeader("Authorization")
		token := strings.Split(tokenString, " ")
		if len(token) != 2 && token[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "request does not contain an access token"})
			c.Abort()
			return
		}
		err := auth.ValidateToken(token[1])
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}

}
