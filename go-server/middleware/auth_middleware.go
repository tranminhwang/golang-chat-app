package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-server/utils"
	"net/http"
	"strings"
	"time"
)

func ValidateToken(c *gin.Context) {
	tokenString := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)

	if len(tokenString) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token is require",
		})
		return
	}

	token, err := utils.ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token isn't valid",
			})
			return
		}
		c.Set("user", map[string]string{
			"email":    claims["email"].(string),
			"username": claims["username"].(string),
			"sub":      claims["sub"].(string),
		})
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token isn't valid",
		})
	}
}
