package news_http

import (
	"fmt"
	"os"

	"github.com/GabrielVilarino/arch-hexagonal/adapter/output/model/response"
	"github.com/GabrielVilarino/arch-hexagonal/application/domain"
	"github.com/GabrielVilarino/arch-hexagonal/configuration/env"
	"github.com/GabrielVilarino/arch-hexagonal/configuration/rest_err"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
)

type newsClient struct {
	client *resty.Client
}

func NewNewsClient() *newsClient {
	return &newsClient{
		client: resty.New().SetBaseURL(os.Getenv("BASE_URL")),
	}
}

func (nc *newsClient) GetNewsPort(newsDomain domain.NewsRequestDomain) (*domain.NewsDomain, *rest_err.RestErr) {

	newsResponse := &response.NewsClientResponse{}

	_, err := nc.client.R().
		SetQueryParams(map[string]string{
			"q":      newsDomain.Subject,
			"from":   newsDomain.From,
			"apiKey": env.GetTokenApi(),
		}).
		SetResult(newsResponse).
		Get("/everything")

	if err != nil {
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("error trying to get news: %s", err.Error()))
	}

	newsResponseDomain := &domain.NewsDomain{}
	copier.Copy(newsResponse, newsResponseDomain)

	return newsResponseDomain, nil
}
