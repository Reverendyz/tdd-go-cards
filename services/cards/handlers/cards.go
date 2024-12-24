package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/reverendyz/tdd-go-cards/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type Card struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func AddCard(c *gin.Context, client *mongo.Client) {
	var card Card
	if err := c.ShouldBindJSON(&card); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := client.Database("cards").Collection("cards")
	if _, err := collection.InsertOne(ctx, card); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to insert card",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Card created successfully",
	})
}

func GetCards(c *gin.Context) {
	var cards []Card
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := db.GetClient()
	coll := client.Database("cards").Collection("cards")
	cursor, err := coll.Find(ctx, Card{})
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Default().Fatalf("error: %v", err.Error())
	}
	cursor.All(ctx, cards)

}
