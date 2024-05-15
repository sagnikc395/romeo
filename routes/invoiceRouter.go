package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sagnikc395/romeo/controllers"
)

//generate invoices for the food ordered by the user

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoices())
	incomingRoutes.PATCH("/invoices/:invoice_id", controller.UpdateInvoice())

}
