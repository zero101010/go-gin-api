package mongodb-learning

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// Create a task class
type Task struct {
	Title string
	Body  string
}

func handleGetTasks(c *gin.Context) {
	var tasks []Task
	var task Task
	task.Title = "Bake some cake"
	task.Body = `- Make a dough 
	- Eat everything before baking 
	- Pretend you never wanted to bake something in the first place`

	tasks = append(tasks, task)
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func createQuickStart(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://igor:95592007i@cluster0.5idqp.mongodb.net",
	))
	if err != nil {
		log.Fatal(err)
	}
	// Criando banco de dados para usarmos nossa collenction
	quickStartDatabase := client.Database("quickStart")
	// Criando as collections
	podcastCollection := quickStartDatabase.Collection("podcast")
	// episodesCollection := quickStartDatabase.Collection("episodes")
	podcastResults, err := podcastCollection.InsertOne(ctx, bson.M{"title": "Study english", "description": "Learn english"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(podcastResults)
}

// func main() {
// 	r := gin.Default()
// 	r.GET("/tasks/", handleGetTasks)
// 	r.GET("/quickstart/", createQuickStart)
// 	r.Run()
// }
