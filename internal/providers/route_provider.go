package providers

import (
	"github.com/d-alejandro/training-level0/internal/use_cases"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

type RouteProvider struct {
	orderModelShowUseCase *use_cases.OrderModelShowUseCase
}

func NewRouteProvider(orderModelShowUseCase *use_cases.OrderModelShowUseCase) *RouteProvider {
	return &RouteProvider{orderModelShowUseCase}
}

func (routeProvider *RouteProvider) Register() {
	router := gin.Default()
	router.Use(gin.Recovery())

	initApiRoutes(router, routeProvider.orderModelShowUseCase)

	port := viper.GetString("HTTP_PORT")

	log.Println("Http server started on : " + port)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func initApiRoutes(router *gin.Engine, orderModelShowUseCase *use_cases.OrderModelShowUseCase) {
	api := router.Group("/api")
	{
		orderModels := api.Group("/order-models")
		{
			orderModels.GET("/:id", orderModelShowUseCase.Show)
		}
	}
}
