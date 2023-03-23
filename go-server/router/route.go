package router

import (
	"github.com/gin-gonic/gin"
	"go-server/internal/user"
	"go-server/middleware"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r = gin.Default()

	authRoutes := r.Group("/api/v1/auth")
	{
		authRoutes.POST("/register", userHandler.CreateUser)
		authRoutes.POST("/login", userHandler.UserLogin)
	}

	messageRoutes := r.Group("/api/v1/messages")
	messageRoutes.Use(middleware.ValidateToken)
	{
		messageRoutes.GET("/", userHandler.GetMessages)
	}

}

func Start(addr string) error {
	return r.Run(addr)
}
