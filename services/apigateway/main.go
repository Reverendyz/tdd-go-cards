package add

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "foo",
	})
}