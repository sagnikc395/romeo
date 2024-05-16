package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sagnikc395/romeo/db"
	"github.com/sagnikc395/romeo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/go-playground/validator.v9"
)

var foodCollection *mongo.Collection = db.OpenCollection(db.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func GetFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		foodId := ctx.Param("food_id")
		var food models.Food
		defer cancel()
		err := foodCollection.FindOne(c, bson.M{"food_id": foodId}).Decode(&food)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching the food item"})
		}
		ctx.JSON(http.StatusOK, food)
	}
}

// adding the food item
func CreateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var c, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu

		if err := ctx.BindJSON(&food); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		validationErr := validate.Struct(food)
		if validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}

		//first find the menu that already exists
		err := menuCollection.FindOne(c, bson.M{"menu_id": food.Menu_id}).Decode(&menu)

		defer cancel()
		if err != nil {
			msg := "menu was not found"
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2)
		food.Price = &num

		result, inserterr := foodCollection.InsertOne(c, food)
		if inserterr != nil {
			msg := "food item was not created"
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()
		ctx.JSON(http.StatusOK, result)
	}

}

func UpdateFood() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func round(num float64) int {
	return int(num)
}

func toFixed(num float64, precision int) float64 {
	return 0.0
}
