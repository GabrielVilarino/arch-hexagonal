package routes

import (
	"github.com/GabrielVilarino/arch-hexagonal/adapter/input/controller"
	"github.com/gin-gonic/gin"
)

func registerV1(group *gin.RouterGroup, newsController *controller.NewsController) {
	v1 := group.Group("/v1")

	// Rotas para Notícias
	news := v1.Group("/news")
	news.GET("/get-news", newsController.GetNews)
}
