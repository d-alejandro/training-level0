package use_cases

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderModelShowUseCase struct {
	cache CacheInterface
}

func NewOrderModelShowUseCase(cache CacheInterface) *OrderModelShowUseCase {
	return &OrderModelShowUseCase{cache}
}

func (orderModelShowUseCase *OrderModelShowUseCase) Show(context *gin.Context) {
	paramId := context.Param("id")

	model, err := orderModelShowUseCase.cache.GetModel(paramId)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response := gin.H{
		"data": model,
	}
	context.JSON(http.StatusOK, response)
}
