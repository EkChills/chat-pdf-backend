package main

import (
	"log"

	"github.com/EkChills/chat-pdf-backend/db"
	"github.com/EkChills/chat-pdf-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.InitDb()

	if err != nil {
		log.Fatal("could not initialize db:", err)
	}

	defer db.DB.Close()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8080")
}