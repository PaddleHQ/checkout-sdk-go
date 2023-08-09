package abc

import (
	"github.com/PaddleHQ/checkout-sdk-go/apm/ideal"
	"github.com/PaddleHQ/checkout-sdk-go/apm/klarna"
	"github.com/PaddleHQ/checkout-sdk-go/apm/sepa"
	"github.com/PaddleHQ/checkout-sdk-go/client"
	"github.com/PaddleHQ/checkout-sdk-go/configuration"
	"github.com/PaddleHQ/checkout-sdk-go/customers"
	"github.com/PaddleHQ/checkout-sdk-go/disputes"
	events "github.com/PaddleHQ/checkout-sdk-go/events/abc"
	"github.com/PaddleHQ/checkout-sdk-go/instruments/abc"
	payments "github.com/PaddleHQ/checkout-sdk-go/payments/abc"
	"github.com/PaddleHQ/checkout-sdk-go/payments/hosted"
	"github.com/PaddleHQ/checkout-sdk-go/payments/links"
	"github.com/PaddleHQ/checkout-sdk-go/reconciliation"
	"github.com/PaddleHQ/checkout-sdk-go/sources"
	"github.com/PaddleHQ/checkout-sdk-go/tokens"
	webhooks "github.com/PaddleHQ/checkout-sdk-go/webhooks/abc"
)

type Api struct {
	Customers      *customers.Client
	Disputes       *disputes.Client
	Events         *events.Client
	Hosted         *hosted.Client
	Instruments    *abc.Client
	Links          *links.Client
	Payments       *payments.Client
	Reconciliation *reconciliation.Client
	Sources        *sources.Client
	Tokens         *tokens.Client
	Webhooks       *webhooks.Client

	Ideal  *ideal.Client
	Klarna *klarna.Client
	Sepa   *sepa.Client
}

func CheckoutApi(configuration *configuration.Configuration) *Api {
	apiClient := buildBaseClient(configuration)

	api := Api{}
	api.Customers = customers.NewClient(configuration, apiClient)
	api.Disputes = disputes.NewClient(configuration, apiClient)
	api.Events = events.NewClient(configuration, apiClient)
	api.Hosted = hosted.NewClient(configuration, apiClient)
	api.Instruments = abc.NewClient(configuration, apiClient)
	api.Links = links.NewClient(configuration, apiClient)
	api.Payments = payments.NewClient(configuration, apiClient)
	api.Reconciliation = reconciliation.NewClient(configuration, apiClient)
	api.Sources = sources.NewClient(configuration, apiClient)
	api.Tokens = tokens.NewClient(configuration, apiClient)
	api.Webhooks = webhooks.NewClient(configuration, apiClient)

	api.Ideal = ideal.NewClient(configuration, apiClient)
	api.Klarna = klarna.NewClient(configuration, apiClient)
	api.Sepa = sepa.NewClient(configuration, apiClient)
	return &api
}

func buildBaseClient(configuration *configuration.Configuration) client.HttpClient {
	return client.NewApiClient(configuration, configuration.Environment.BaseUri())
}
