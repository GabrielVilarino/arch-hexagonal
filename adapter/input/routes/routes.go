package routes

import (
	"github.com/GabrielVilarino/arch-hexagonal/adapter/input/controller"
	"github.com/GabrielVilarino/arch-hexagonal/adapter/output/news_http"
	"github.com/GabrielVilarino/arch-hexagonal/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newsClient := news_http.NewNewsClient()

	newsService := service.NewNewsService(newsClient)

	newsController := controller.NewNewsController(newsService)

	api := r.Group("/api")

	registerV1(api, newsController)
}
