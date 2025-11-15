package output

import (
	"github.com/GabrielVilarino/arch-hexagonal/application/domain"
	"github.com/GabrielVilarino/arch-hexagonal/configuration/rest_err"
)

type NewsPort interface {
	GetNewsPort(domain.NewsRequestDomain) (*domain.NewsDomain, *rest_err.RestErr)
}
