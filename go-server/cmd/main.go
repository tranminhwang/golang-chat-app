package main

import (
	"github.com/joho/godotenv"
	"go-server/database"
	"go-server/internal/user"
	"go-server/router"
	"go-server/ws"
	"log"
)

func main() {
	godotenv.Load()
	dbConn, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}

	userRepository := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}
