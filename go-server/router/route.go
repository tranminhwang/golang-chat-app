package router

import (
	"github.com/gin-gonic/gin"
	"go-server/internal/user"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r = gin.Default()

	r.POST("/api/v1/auth/register", userHandler.CreateUser)
	// middleware.ValidateToken
	r.POST("/api/v1/auth/login", userHandler.UserLogin)
}

func Start(addr string) error {
	return r.Run(addr)
}
