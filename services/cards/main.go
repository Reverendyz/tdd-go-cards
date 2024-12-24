package main

import (
	"fmt"
	"log"
	"reverendyz/tdd-go-cards/services/cards/handlers"

	"github.com/gin-gonic/gin"
)

var (
	host string
	port string
)

func init() {

}

func main() {
	router := gin.Default()
	router.POST("/cards", handlers.AddCard)
	router.GET("/cards", handlers.GetCards)
	// router.PUT("/cards/:id", updateCardHandler)
	// router.DELETE("/cards/:id", deleteCardHandler)

	if err := router.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
