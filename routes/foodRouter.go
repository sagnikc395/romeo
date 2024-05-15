package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sagnikc395/romeo/controllers"
)

//router for food and menu options

func FoodRoutes(incomingRoutes *gin.Engine) {
	//all food items
	incomingRoutes.GET("/foods", controller.GetFoods())
	//particular food item
	incomingRoutes.GET("/foods/:food_id", controller.GetFood())
	//add new food item
	incomingRoutes.POST("/foods/", controller.CreateFood())
	//update food item
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
}
