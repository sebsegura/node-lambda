package handler

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"seb7887/go-trx-metrics/internal/processor"
	"seb7887/go-trx-metrics/pkg/models"
)

type Response = events.APIGatewayProxyResponse

type Handler interface {
	AddMetrics(ctx context.Context, req models.Request) (Response, error)
}

type handler struct {
	processor processor.Processor
}

func New(p processor.Processor) Handler {
	return &handler{
		processor: p,
	}
}

func (h *handler) AddMetrics(_ context.Context, req models.Request) (Response, error) {
	metric, err := h.processor.Process(req)
	if err != nil {
		return Response{
			StatusCode: http.StatusBadRequest,
			Body:       "{\"message\": \"REJECTED\"}",
		}, nil
	}

	return Response{
		StatusCode: http.StatusOK,
		Body:       metric.ToJsonStr(),
	}, nil
}
