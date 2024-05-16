package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sagnikc395/romeo/db"
	"github.com/sagnikc395/romeo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = db.OpenCollection(db.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//return all menus using Find
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		result, err := menuCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing the menu items"})
		}

		var allMenus []bson.M
		if err = result.All(c, &allMenus); err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, allMenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		menuId := ctx.Param("menu_id")
		var menu models.Menu

		err := foodCollection.FindOne(c, bson.M{"menu_id": menuId}).Decode(&menu)

		defer cancel()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the menu"})

		}
		ctx.JSON(http.StatusOK, menu)
	}
}

func CreateMenus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var menu models.Menu
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		if err := ctx.BindJSON(&menu); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		//validation and validation error handler

		validationErr := validate.Struct(menu)

		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
