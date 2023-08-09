package abc

import (
	"net/http"

	"github.com/PaddleHQ/checkout-sdk-go/configuration"
)

type CheckoutPreviousSdkBuilder struct {
	configuration.StaticKeysBuilder
}

func (b *CheckoutPreviousSdkBuilder) WithEnvironment(environment configuration.Environment) *CheckoutPreviousSdkBuilder {
	b.Environment = environment
	return b
}

func (b *CheckoutPreviousSdkBuilder) WithHttpClient(client *http.Client) *CheckoutPreviousSdkBuilder {
	b.HttpClient = client
	return b
}

func (b *CheckoutPreviousSdkBuilder) WithLogger(logger configuration.StdLogger) *CheckoutPreviousSdkBuilder {
	b.Logger = logger
	return b
}

func (b *CheckoutPreviousSdkBuilder) WithPublicKey(publicKey string) *CheckoutPreviousSdkBuilder {
	b.PublicKey = publicKey
	return b
}

func (b *CheckoutPreviousSdkBuilder) WithSecretKey(secretKey string) *CheckoutPreviousSdkBuilder {
	b.SecretKey = secretKey
	return b
}

func (b *CheckoutPreviousSdkBuilder) Build() (*Api, error) {
	err := b.ValidateSecretKey(configuration.PreviousSecretKeyPattern)
	if err != nil {
		return nil, err
	}

	err = b.ValidatePublicKey(configuration.PreviousPublicKeyPattern)
	if err != nil {
		return nil, err
	}

	sdkCredentials := configuration.NewPreviousKeysSdkCredentials(b.SecretKey, b.PublicKey)
	newConfiguration := configuration.NewConfiguration(sdkCredentials, b.Environment, b.HttpClient, b.Logger)

	return CheckoutApi(newConfiguration), nil
}
