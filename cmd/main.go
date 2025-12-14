package main

import (
	"log"

	"github.com/Tommy-jeff/first_go_crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/// função main é o ponto de entrada da nossa aplicação
func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Error loading .env file -> ", err)
	}

	/// inicializamos os nossos endpoints
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	/// inicializamos nossa aplicação e começamos a ouvir na porta 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}