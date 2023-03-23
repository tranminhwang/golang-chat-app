package router

import (
	"github.com/gin-gonic/gin"
	"go-server/internal/user"
	"go-server/middleware"
	"go-server/ws"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
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

	wsRoutes := r.Group("/ws")
	wsRoutes.Use(middleware.ValidateToken)
	{
		wsRoutes.POST("/create_room", wsHandler.CreateRoom)
		wsRoutes.GET("/join_room/:roomId", wsHandler.JoinRoom)
		wsRoutes.GET("/get_rooms", wsHandler.GetRooms)
		wsRoutes.GET("/get_clients/:roomId", wsHandler.GetClients)
	}

}

func Start(addr string) error {
	return r.Run(addr)
}
