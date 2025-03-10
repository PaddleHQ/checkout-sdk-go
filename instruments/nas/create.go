package nas

import (
	"encoding/json"

	"github.com/PaddleHQ/checkout-sdk-go/common"
)

type (
	CreateInstrumentRequest interface{}

	createBankAccountInstrumentRequest struct {
		Type                common.InstrumentType            `json:"type" binding:"required"`
		AccountType         common.AccountType               `json:"account_type,omitempty"`
		AccountNumber       string                           `json:"account_number,omitempty"`
		BankCode            string                           `json:"bank_code,omitempty"`
		BranchCode          string                           `json:"branch_code,omitempty"`
		Iban                string                           `json:"iban,omitempty"`
		Bban                string                           `json:"bban,omitempty"`
		SwiftBic            string                           `json:"swift_bic,omitempty"`
		Currency            common.Currency                  `json:"currency,omitempty"`
		Country             common.Country                   `json:"country,omitempty"`
		ProcessingChannelId string                           `json:"processing_channel_id,omitempty"`
		AccountHolder       *common.AccountHolder            `json:"account_holder,omitempty"`
		BankDetails         *common.BankDetails              `json:"bank,omitempty"`
		Customer            *CreateCustomerInstrumentRequest `json:"customer,omitempty"`
	}

	createTokenInstrumentRequest struct {
		Type          common.InstrumentType            `json:"type" binding:"required"`
		Token         string                           `json:"token" binding:"required"`
		AccountHolder *common.AccountHolder            `json:"account_holder" binding:"required"`
		Customer      *CreateCustomerInstrumentRequest `json:"customer,omitempty"`
	}

	createSepaInstrumentRequest struct {
		Type           common.InstrumentType `json:"type" binding:"required"`
		InstrumentData *InstrumentData       `json:"instrument_data,omitempty"`
		AccountHolder  *common.AccountHolder `json:"account_holder" binding:"required"`
	}
)

func NewCreateBankAccountInstrumentRequest() *createBankAccountInstrumentRequest {
	return &createBankAccountInstrumentRequest{
		Type: common.BankAccount,
	}
}

func NewCreateTokenInstrumentRequest() *createTokenInstrumentRequest {
	return &createTokenInstrumentRequest{
		Type: common.Token,
	}
}

func NewCreateSepaInstrumentRequest() *createSepaInstrumentRequest {
	return &createSepaInstrumentRequest{
		Type: common.Sepa,
	}
}

type (
	CreateInstrumentResponse struct {
		HttpMetadata                        common.HttpMetadata
		CreateBankAccountInstrumentResponse *CreateBankAccountInstrumentResponse
		CreateTokenInstrumentResponse       *CreateTokenInstrumentResponse
		CreateSepaInstrumentResponse        *CreateSepaInstrumentResponse
		AlternativeResponse                 *common.AlternativeResponse
	}

	CreateBankAccountInstrumentResponse struct {
		Type common.InstrumentType `json:"type" binding:"required"`
		// common
		Id               string                   `json:"id,omitempty"`
		Fingerprint      string                   `json:"fingerprint,omitempty"`
		CustomerResponse *common.CustomerResponse `json:"customer,omitempty"`
		// specific
		BankDetails   *common.BankDetails `json:"bank,omitempty"`
		SwiftBic      string              `json:"swift_bic,omitempty"`
		AccountNumber string              `json:"account_number,omitempty"`
		BankCode      string              `json:"bank_code,omitempty"`
		Iban          string              `json:"iban,omitempty"`
	}

	CreateTokenInstrumentResponse struct {
		Type common.InstrumentType `json:"type" binding:"required"`
		// common
		Id               string                   `json:"id,omitempty"`
		Fingerprint      string                   `json:"fingerprint,omitempty"`
		CustomerResponse *common.CustomerResponse `json:"customer,omitempty"`
		// specific
		ExpiryMonth   int                 `json:"expiry_month,omitempty"`
		ExpiryYear    int                 `json:"expiry_year,omitempty"`
		Scheme        string              `json:"scheme,omitempty"`
		SchemeLocal   string              `json:"scheme_local,omitempty"`
		Last4         string              `json:"last4,omitempty"`
		Bin           string              `json:"bin,omitempty"`
		CardType      common.CardType     `json:"card_type,omitempty"`
		CardCategory  common.CardCategory `json:"card_category,omitempty"`
		Issuer        string              `json:"issuer,omitempty"`
		IssuerCountry common.Country      `json:"issuer_country,omitempty"`
		ProductId     string              `json:"product_id,omitempty"`
		ProductType   string              `json:"product_type,omitempty"`
	}

	CreateSepaInstrumentResponse struct {
		Type common.InstrumentType `json:"type" binding:"required"`
		// common
		Id          string `json:"id,omitempty"`
		Fingerprint string `json:"fingerprint,omitempty"`
	}
)

func (s *CreateInstrumentResponse) UnmarshalJSON(data []byte) error {
	var typeMapping common.TypeMapping
	if err := json.Unmarshal(data, &typeMapping); err != nil {
		return err
	}

	switch typeMapping.Type {
	case string(common.BankAccount):
		var response CreateBankAccountInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.CreateBankAccountInstrumentResponse = &response
	case string(common.Card):
		var response CreateTokenInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.CreateTokenInstrumentResponse = &response
	case string(common.Sepa):
		var response CreateSepaInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.CreateSepaInstrumentResponse = &response
	default:
		var response common.AlternativeResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.AlternativeResponse = &response
	}

	return nil
}
