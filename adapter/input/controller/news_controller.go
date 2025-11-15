package controller

import (
	"net/http"

	"github.com/GabrielVilarino/arch-hexagonal/adapter/input/model/request"
	"github.com/GabrielVilarino/arch-hexagonal/application/domain"
	"github.com/GabrielVilarino/arch-hexagonal/application/port/input"
	"github.com/GabrielVilarino/arch-hexagonal/configuration/logger"
	"github.com/GabrielVilarino/arch-hexagonal/configuration/validation"
	"github.com/gin-gonic/gin"
)

type NewsController struct {
	newsUseCase input.NewsUseCase
}

func NewNewsController(
	newsUseCase input.NewsUseCase,
) *NewsController {
	return &NewsController{newsUseCase}
}

func (nc *NewsController) GetNews(c *gin.Context) {
	newsRequest := request.NewsRequest{}

	if err := c.ShouldBindQuery(&newsRequest); err != nil {
		logger.Error("Error trying to validate fields from request", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	newsDomain := domain.NewsRequestDomain{Subject: newsRequest.Subject, From: newsRequest.From.Format("2006-01-02")}

	newsResponseDomain, err := nc.newsUseCase.GetNewsService(newsDomain)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, newsResponseDomain)
}
