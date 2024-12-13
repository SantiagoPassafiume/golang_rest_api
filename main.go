package main

import (
	"github.com/SantiagoPassafiume/golang_rest_api/db"
	"github.com/SantiagoPassafiume/golang_rest_api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}
