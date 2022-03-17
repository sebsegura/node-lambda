package handler

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
	"seb7887/go-trx-processor/internal/processor"
	"seb7887/go-trx-processor/pkg/models"
)

type Request = events.APIGatewayProxyRequest
type Response = events.APIGatewayProxyResponse

type Handler interface {
	ProcessTrx(ctx context.Context, req Request) (Response, error)
}

type handler struct {
	processor processor.Processor
}

func New(p processor.Processor) Handler {
	return &handler{
		processor: p,
	}
}

func (h *handler) ProcessTrx(_ context.Context, req Request) (Response, error) {
	log.Println(req)
	var r models.Request
	_ = json.Unmarshal([]byte(req.Body), &r)
	if err := h.processor.Process(r); err != nil {
		log.Println(err.Error())
		return Response{
			StatusCode: http.StatusBadRequest,
			Body:       "{\"status\":\"REJECTED\"}",
		}, nil
	}

	return Response{
		StatusCode: http.StatusOK,
		Body:       "{\"status\":\"PROCESSED\"}",
	}, nil
}
