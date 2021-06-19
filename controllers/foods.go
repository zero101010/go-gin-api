package controllers

import (
	"context"
	"fmt"
	"go-gin-api/database"
	"go-gin-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func CreateFood(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var food models.Food
	if err := c.BindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	result, insertErr := foodCollection.InsertOne(ctx, food)
	if insertErr != nil {
		msg := fmt.Sprintf("Food item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()

	//return the id of the created object to the frontend
	c.JSON(http.StatusOK, result)

}

func GetFoods(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var foods []models.Food
	result, insertErr := foodCollection.Find(ctx, bson.M{})
	if insertErr != nil {
		msg := fmt.Sprintf("Food item was not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()
	for result.Next(ctx) {
		//Create a value into which the single document can be decoded
		var food models.Food
		result.Decode(&food)
		foods = append(foods, food)

	}
	c.JSON(http.StatusOK, foods)
}

func GetFood(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	fmt.Println(id)
	var food models.Food
	err := foodCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&food)
	fmt.Println(food)
	if err != nil {
		msg := fmt.Sprintf("Food item was not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, food)
}
