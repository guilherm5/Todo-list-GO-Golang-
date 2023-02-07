package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crud-gingonic-mvc/autentication"
	"github.com/guilherm5/crud-gingonic-mvc/database"
	routes "github.com/guilherm5/crud-gingonic-mvc/routes"
	"github.com/joho/godotenv"
)

func main() {
	database.Init()
	autentication.Config()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	port := os.Getenv("PORT")
	router := gin.New()

	routes.TodoList(router)
	routes.Autentication(router)

	router.Run(":" + port)
}
