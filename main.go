package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ridhohrnf/task-5-vix-btpns-Muhammad-Ridho-Heranof/app"
	"github.com/ridhohrnf/task-5-vix-btpns-Muhammad-Ridho-Heranof/routes"
)

func main() {
	router := gin.Default()

	db := app.SetupDatabase()
	defer db.Close()

	routes.SetupUserRoutes(router, db)
	routes.SetupPhotoRoutes(router, db)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
