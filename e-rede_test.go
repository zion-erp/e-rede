package erede

import (
	"crypto/rand"
	"fmt"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/zion-erp/e-rede/enums"
	"github.com/zion-erp/e-rede/models"
)

type TestERedeClient struct{}

func (s *TestERedeClient) NewTestClient() *ERedeClient {
	c := &Config{}
	c.Init("85587178", "1db5cdf134774d4bb20719ea21ff7d8d", "v1", true)
	l := hclog.Default()
	return NewClient(*c, l)
}

func (s *TestERedeClient) NewRandomId() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", b)
}

func (s *TestERedeClient) TestCreateTransaction() *models.TransactionResponse {
	reference := s.NewRandomId()
	cl := s.NewTestClient()
	tr := &models.Transaction{
		Capture:                false,
		Kind:                   enums.TransactionKind_Credit,
		Reference:              reference,
		Amount:                 2000,
		Installments:           2,
		CardHolderName:         "John Snow",
		CardNumber:             "5448280000000007",
		ExpirationMonth:        "01",
		ExpirationYear:         "2028",
		SecurityCode:           "123",
		SoftDescriptor:         reference,
		Subscription:           false,
		Origin:                 enums.TransactionOrigin_OriginERede,
		DistributorAffiliation: 0,
		// BrandTid:               "string",
		StorageCard: enums.TransactionCardStorage_CredentialNotStored,
		TransactionCredentials: &models.TransactionCredentials{
			CredentialId: "01",
		},
	}
	transaction, resp := cl.CreateTransaction(tr)

	if resp.Error != nil {
		cl.log.Debug(resp.Error.Error())
	}

	fmt.Printf("%+v\n", transaction)
	return transaction
}

func (s *TestERedeClient) TestCreateTransactionAirlines() {
	reference := s.NewRandomId()
	cl := s.NewTestClient()
	tr := &models.Transaction{
		Capture:         true,
		Kind:            enums.TransactionKind_Credit,
		Reference:       reference,
		Amount:          2000,
		CardHolderName:  "John Snow",
		CardNumber:      "5448280000000007",
		ExpirationMonth: "12",
		ExpirationYear:  "2028",
		SecurityCode:    "235",
		SoftDescriptor:  reference,
		Origin:          enums.TransactionOrigin_OriginERede,
		Iata: &models.Iata{
			Code:         101010,
			DepartureTax: 100,
		},
	}
	transaction, resp := cl.CreateTransaction(tr)

	if resp.Error != nil {
		cl.log.Debug(resp.Error.Error())
	}

	fmt.Printf("%+v\n", transaction)
}

func (s *TestERedeClient) TestCreateTransactionZeroDollar() {
	reference := s.NewRandomId()
	cl := s.NewTestClient()
	tr := &models.Transaction{
		Capture:         true,
		Kind:            enums.TransactionKind_Credit,
		Reference:       reference,
		Amount:          0,
		CardHolderName:  "John Snow",
		CardNumber:      "5448280000000007",
		ExpirationMonth: "12",
		ExpirationYear:  "2028",
		SecurityCode:    "235",
		SoftDescriptor:  reference,
		Origin:          enums.TransactionOrigin_OriginERede,
	}
	transaction, resp := cl.CreateTransaction(tr)

	if resp.Error != nil {
		cl.log.Debug(resp.Error.Error())
	}

	fmt.Printf("%+v\n", transaction)
}

func (s *TestERedeClient) TestGetTransaction(tid string) {
	cl := s.NewTestClient()
	transaction, resp := cl.GetTransactionByTid(tid)

	if resp.Error != nil {
		cl.log.Debug(resp.Error.Error())
	}

	fmt.Printf("%+v\n", transaction)
}

func (s *TestERedeClient) TestGetTransactionByReference(reference string) {
	cl := s.NewTestClient()
	transaction, resp := cl.GetTransactionByReference(reference)

	if resp.Error != nil {
		cl.log.Debug(resp.Error.Error())
	}

	fmt.Printf("%+v\n", transaction)
}

func (s *TestERedeClient) TestCaptureTransaction(tid string, amount float64) {
	cl := s.NewTestClient()
	transaction, resp := cl.CaptureTransaction(tid, amount)

	if resp.Error != nil {
		cl.log.Debug(resp.Error.Error())
	}

	fmt.Printf("%+v\n", transaction)
}

func (s *TestERedeClient) TestTokenizeCard() {
	cl := s.NewTestClient()
	card, resp := cl.CreateTokenization(&models.CreateTokenization{
		CardholderName:  "John Snow",
		CardNumber:      "4895370010000005",
		Email:           "john_snow@gmail.com",
		ExpirationMonth: "02",
		ExpirationYear:  "2028",
		SecurityCode:    "235",
		StorageCard:     enums.TransactionCardStorage_CredentialNotStored,
	})

	if resp.Error != nil {
		cl.log.Debug(resp.Error.Error())
	}

	fmt.Printf("%+v\n", card)
}
