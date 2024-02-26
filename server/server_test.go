package main

import (
	"context"
	"testing"

	pb "currency_convert/proto"

	"github.com/stretchr/testify/assert"
)

func TestConvertCurrencyToOtherCurrency(t *testing.T) {
	server := server{}

	req := &pb.CurrencyConvertRequest{
		Amount:       10,
		FromCurrency: "RUPEE",
		ToCurrency:   "DOLLAR",
	}
	expectedRes := &pb.CurrencyConvertResponse{
		Amount: 800,
	}

	res, err := server.Convert(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Amount)
	assert.Equal(t, expectedRes, res)
}

func TestConvertCurrencyToSameCurrency(t *testing.T) {
	server := server{}

	req := &pb.CurrencyConvertRequest{
		Amount:       10,
		FromCurrency: "RUPEE",
		ToCurrency:   "RUPEE",
	}
	expectedRes := &pb.CurrencyConvertResponse{
		Amount: 10,
	}

	res, err := server.Convert(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Amount)
	assert.Equal(t, expectedRes, res)
}
