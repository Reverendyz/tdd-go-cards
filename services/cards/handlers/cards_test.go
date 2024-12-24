package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestAddCard(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("add valid card", func(mt *mtest.T) {

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/cards", handlers.AddCard)

		card := handlers.Card{
			Title:       "Test Card",
			Description: "Test Description",
		}
		jsonData, _ := json.Marshal(card)
		req := httptest.NewRequest(http.MethodPost, "/cards", bytes.NewReader(jsonData))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		mt.AddMockResponses(mtest.CreateSuccessResponse())
		db.GetClient = func() *mongo.Client {
			return mt.Client
		}

		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		var response map[string]interface{}
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "message")
		assert.Equal(t, "Card created successfully", response["message"])
	})

	mt.Run("add invalid card", func(mt *mtest.T) {

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/cards", handlers.AddCard)

		invalidCard := map[string]string{
			"Title": "",
		}
		jsonData, _ := json.Marshal(invalidCard)
		req := httptest.NewRequest(http.MethodPost, "/cards", bytes.NewReader(jsonData))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		var response map[string]interface{}
		err := json.Unmarshal(resp.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Contains(t, response, "error")
	})
}
