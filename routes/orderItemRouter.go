package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sagnikc395/romeo/controllers"
)

// table management
func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:order_item_id", controller.GetOrderItem())
	//get all the order items inside the router
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItems())
	incomingRoutes.PATCH("/orderItems/:order_item_id", controller.UpdateOrderItem())
}
