package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-server/internal/user"
	"go-server/middleware"
	"go-server/ws"
	"os"
	"time"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("CLIENT_URL")},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == os.Getenv("CLIENT_URL")
		},
		MaxAge: time.Hour * 12,
	}))

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
	// wsRoutes.Use(middleware.ValidateToken)
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
