package handlers

import (
	"context"
	"log"
	"net/http"
	"reverendyz/tdd-go-cards/pkg/db"
	"time"

	"github.com/gin-gonic/gin"
)

type Card struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func AddCard(c *gin.Context) {
	var card Card
	if err := c.ShouldBindJSON(&card); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := db.GetClient()

	client.Database("cards").Collection("cards").InsertOne(ctx, card)

	client.Disconnect(ctx)
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
		log.Default().Fatalf("error: ", err.Error())
	}
	cursor.All(ctx, cards)

}
