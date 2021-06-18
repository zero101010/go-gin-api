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
	//generate new ID for the object to be created
	food.ID = primitive.NewObjectID()

	// assign the the auto generated ID to the primary key attribute
	food.Food_id = food.ID.Hex()
	//insert the newly created object into mongodb
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

// func GetFoods(c *gin.Context) {
// 	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 	var foods []models.Food
// 	result, insertErr := foodCollection.Find(ctx, bson.M{}).All(&foods)
// 	fmt.Println(result)
// 	if insertErr != nil {
// 		msg := fmt.Sprintf("Food item was not found")
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 		return
// 	}
// 	fmt.Println(result)
// 	c.JSON(http.StatusOK, result)
// 	defer cancel()
// }
