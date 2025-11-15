package input

import (
	"github.com/GabrielVilarino/arch-hexagonal/application/domain"
	"github.com/GabrielVilarino/arch-hexagonal/configuration/rest_err"
)

type NewsUseCase interface {
	GetNewsService(
		domain.NewsRequestDomain,
	) (*domain.NewsDomain, *rest_err.RestErr)
}
