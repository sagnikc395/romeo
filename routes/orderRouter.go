package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sagnikc395/romeo/controllers"
)

// order management of food id and current status
func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order", controller.GetOrders())
	incomingRoutes.GET("/order/:order_id", controller.GetOrder())
	incomingRoutes.POST("/orders", controller.CreateOrders())
	incomingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())
}
