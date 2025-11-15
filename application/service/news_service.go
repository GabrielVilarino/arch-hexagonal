package service

import (
	"fmt"

	"github.com/GabrielVilarino/arch-hexagonal/application/domain"
	"github.com/GabrielVilarino/arch-hexagonal/application/port/output"
	"github.com/GabrielVilarino/arch-hexagonal/configuration/logger"
	"github.com/GabrielVilarino/arch-hexagonal/configuration/rest_err"
)

type newsService struct {
	newsPort output.NewsPort
}

func NewNewsService(newsPort output.NewsPort) *newsService {
	return &newsService{newsPort}
}

func (ns *newsService) GetNewsService(
	req domain.NewsRequestDomain,
) (*domain.NewsDomain, *rest_err.RestErr) {
	logger.Info(fmt.Sprintf("Iniciando a função getNewsService, assunto=%s, desde=%s", req.Subject, req.From))

	newsDomainResponse, err := ns.newsPort.GetNewsPort(req)

	return newsDomainResponse, err
}
