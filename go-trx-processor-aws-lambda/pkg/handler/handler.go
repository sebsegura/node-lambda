package handler

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"seb7887/go-trx-processor/internal/processor"
	"seb7887/go-trx-processor/pkg/models"
)

type Response = events.APIGatewayProxyResponse

type Handler interface {
	ProcessTrx(ctx context.Context, req models.Request) (Response, error)
}

type handler struct {
	processor processor.Processor
}

func New(p processor.Processor) Handler {
	return &handler{
		processor: p,
	}
}

func (h *handler) ProcessTrx(_ context.Context, req models.Request) (Response, error) {
	if err := h.processor.Process(req); err != nil {
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
