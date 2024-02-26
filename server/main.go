package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "currency_convert/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCurrencyConvertServiceServer
}

type Currency string

const (
	RUPEE  Currency = "RUPEE"
	DOLLAR Currency = "DOLLAR"
)

var StringToCurrencyMap = map[string]Currency{
	"RUPEE":  RUPEE,
	"DOLLAR": DOLLAR,
}

var currencyConvertRate = map[Currency]float32{
	RUPEE:  1,
	DOLLAR: 80,
}

func (c Currency) convertToBase(amount float32) float32 {
	return amount / currencyConvertRate[c]
}

func (c Currency) convertFromBase(amount float32) float32 {
	return amount * currencyConvertRate[c]
}

func (s *server) Convert(ctx context.Context, req *pb.CurrencyConvertRequest) (*pb.CurrencyConvertResponse, error) {
	var fromCurrency Currency
	if currency, ok := StringToCurrencyMap[req.FromCurrency]; ok {
		fromCurrency = currency
	} else {
		return &pb.CurrencyConvertResponse{}, fmt.Errorf("unsupported currency: %s", req.FromCurrency)
	}

	var toCurrency Currency
	if currency, ok := StringToCurrencyMap[req.ToCurrency]; ok {
		toCurrency = currency
	} else {
		return &pb.CurrencyConvertResponse{}, fmt.Errorf("unsupported currency: %s", req.ToCurrency)
	}

	fmt.Println(fromCurrency.convertToBase(req.Amount))
	fmt.Println(toCurrency.convertFromBase(fromCurrency.convertToBase(req.Amount)))
	convertedAmount := toCurrency.convertFromBase(fromCurrency.convertToBase(req.Amount))
	return &pb.CurrencyConvertResponse{Amount: convertedAmount}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterCurrencyConvertServiceServer(srv, &server{})

	log.Println("Starting gRPC server on port 50051...")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
