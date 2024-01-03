package erede

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/zion-erp/e-rede/models"
)

func (s *ERedeClient) GetTransactionByTid(tid string) (*models.TransactionResponse, *Response) {
	url := fmt.Sprintf("%s/transactions/%s", s.Config.ApiUrl, tid)

	response := s.Get(url)

	if !response.Ok() {
		return nil, response
	}

	res, err := ByteArrayToStruct[models.TransactionResponse](response.Payload)
	if err != nil {
		response.Error = err
	}

	return res, response
}

func (s *ERedeClient) GetTransactionByReference(reference string) (*models.TransactionResponse, *Response) {
	url := fmt.Sprintf("%s/transactions?reference=%s", s.Config.ApiUrl, reference)

	response := s.Get(url)

	if !response.Ok() {
		return nil, response
	}

	res, err := ByteArrayToStruct[models.TransactionResponse](response.Payload)
	if err != nil {
		response.Error = err
	}

	return res, response
}

func (c *ERedeClient) CreateTransaction(transaction *models.Transaction) (*models.TransactionResponse, *Response) {
	url := fmt.Sprintf("%s/transactions/%s", c.Config.ApiUrl, transaction.Tid)

	response := c.Post(url, transaction)

	if !response.OkCreated() {
		return nil, response
	}

	res, err := ByteArrayToStruct[models.TransactionResponse](response.Payload)
	if err != nil {
		response.Error = err
	}

	return res, response
}

func (c *ERedeClient) CaptureTransaction(tid string, amount float64) (*models.TransactionResponse, *Response) {
	url := fmt.Sprintf("%s/transactions/%s", c.Config.ApiUrl, tid)

	cp := &models.Capture{
		Amount: uint64(amount * 100),
	}

	response := c.Put(url, cp)

	if !response.OkCreated() {
		return nil, response
	}

	res, err := ByteArrayToStruct[models.TransactionResponse](response.Payload)
	if err != nil {
		response.Error = err
	}

	return res, response
}

func (s *ERedeClient) GetCancellationsByTid(tid string) (*models.TransactionResponse, *Response) {
	url := fmt.Sprintf("%s/transactions/%s/refunds", s.Config.ApiUrl, tid)

	response := s.Get(url)

	if !response.Ok() {
		return nil, response
	}

	res, err := ByteArrayToStruct[models.TransactionResponse](response.Payload)
	if err != nil {
		response.Error = err
	}

	return res, response
}

func (s *ERedeClient) GetCancellationByRefundId(tid string, refundId string) (*models.Refund, *Response) {
	url := fmt.Sprintf("%s/transactions/%s/refunds/%s", s.Config.ApiUrl, tid, refundId)

	response := s.Get(url)

	if !response.Ok() {
		return nil, response
	}

	res, err := ByteArrayToStruct[models.Refund](response.Payload)
	if err != nil {
		response.Error = err
	}

	return res, response
}

func (c *ERedeClient) CancelTransaction(tid string, transaction *models.Transaction) (*models.Refund, *Response) {
	url := fmt.Sprintf("%s/transactions/%s/refunds", c.Config.ApiUrl, tid)

	response := c.Post(url, transaction)

	if !response.OkCreated() {
		return nil, response
	}

	res, err := ByteArrayToStruct[models.Refund](response.Payload)
	if err != nil {
		response.Error = err
	}

	return res, response
}

func (s *ERedeClient) GetTokenizationById(tokenizationId *uuid.UUID) (*models.QueryTokenizationResponse, *Response) {
	url := fmt.Sprintf("%s/tokenization/%s", s.Config.ApiUrl, tokenizationId.String())

	response := s.Get(url)

	if !response.Ok() {
		return nil, response
	}

	res, err := ByteArrayToStruct[models.QueryTokenizationResponse](response.Payload)
	if err != nil {
		response.Error = err
	}

	return res, response
}

func (s *ERedeClient) GetCryptogramByTokenizationId(tokenizationId *uuid.UUID, subscription bool) (*models.QueryTokenizationCryptogramResponse, *Response) {
	url := fmt.Sprintf("%s/cryptogram/%s", s.Config.ApiUrl, tokenizationId.String())

	data := map[string]bool{"subscription": subscription}
	response := s.Post(url, data)

	if !response.Ok() {
		return nil, response
	}

	res, err := ByteArrayToStruct[models.QueryTokenizationCryptogramResponse](response.Payload)
	if err != nil {
		response.Error = err
	}

	return res, response
}

func (c *ERedeClient) CreateTokenization(tokenization *models.CreateTokenization) (*models.CreateTokenizationResponse, *Response) {
	url := fmt.Sprintf("%s/tokenization", c.Config.ApiUrl)

	response := c.Post(url, tokenization)

	if !response.OkCreated() {
		return nil, response
	}

	res, err := ByteArrayToStruct[models.CreateTokenizationResponse](response.Payload)
	if err != nil {
		response.Error = err
	}

	return res, response
}
