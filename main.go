package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sagnikc395/romeo/db"
	"github.com/sagnikc395/romeo/middleware"
	"github.com/sagnikc395/romeo/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = db.OpenCollection(db.Client, "food")

func main() {
	//connect to the database
	port := os.Getenv("PORT")

	if port == "" {
		//default port
		port = "8080"
	}

	//using gin for routing
	router := gin.New()
	router.Use(gin.Logger())
	//register routes
	routes.UserRoutes(router)
	//add auth middlewar
	router.Use(middleware.Authentication())

	//assign all routes to the router
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	//start the server
	router.Run(":" + port)
}
