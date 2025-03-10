package financial

import (
	"github.com/PaddleHQ/checkout-sdk-go/client"
	"github.com/PaddleHQ/checkout-sdk-go/common"
	"github.com/PaddleHQ/checkout-sdk-go/configuration"
)

type Client struct {
	configuration *configuration.Configuration
	apiClient     client.HttpClient
}

func NewClient(configuration *configuration.Configuration, apiClient client.HttpClient) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
	}
}

func (c *Client) GetFinancialActions(query QueryFilter) (*QueryResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(FinancialActionsPath), query)
	if err != nil {
		return nil, err
	}

	var response QueryResponse
	err = c.apiClient.Get(
		url,
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
