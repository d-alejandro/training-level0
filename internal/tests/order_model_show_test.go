package tests

import (
	"encoding/json"
	"github.com/d-alejandro/training-level0/internal/bootstrap"
	"github.com/d-alejandro/training-level0/internal/models"
	"github.com/d-alejandro/training-level0/internal/resources"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const idPrefix = "api-test"

func TestOrderModelShow(t *testing.T) {
	bootstrap.Boot()

	expectedMessage := resources.GetModelJSON(idPrefix)
	addMessageToCache(expectedMessage)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/order-models/b563feb7b2b84b6test"+idPrefix, nil)
	bootstrap.Router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)

	expectedString := `{"data":` + expectedMessage + `}`
	actualString := response.Body.String()
	assert.JSONEq(t, expectedString, actualString)
}

func addMessageToCache(message string) {
	var model models.Model

	if err := json.Unmarshal([]byte(message), &model); err != nil {
		log.Fatal(err)
	}

	if err := bootstrap.Cache.SetModel(model.OrderUID, &model); err != nil {
		log.Fatal(err)
	}
}
