package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func Authorize(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}
	authSplit := strings.Split(authHeader, "Bearer ")
	if len(authSplit) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
		return
	}
	userId, _ := strconv.ParseInt(authSplit[1], 10, 64)
	c.Set("userId", userId)
	c.Next()
}
