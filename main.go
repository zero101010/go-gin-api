package main

import (
	"go-gin-api/controllers"
	"go-gin-api/database"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// Create food Class

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.Default()
	// router.Use(gin.Logger())

	router.POST("/foods", controllers.CreateFood)
	router.GET("/foods", controllers.GetFoods)
	router.GET("/foods/:id", controllers.GetFood)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, bson.M{"Application": "Food teste", "Status": "Up"})
	})
	router.Run(":" + port)
}
