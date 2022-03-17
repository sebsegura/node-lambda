package handler

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"seb7887/go-trx-metrics/internal/processor"
	"seb7887/go-trx-metrics/pkg/models"
)

type Request = events.APIGatewayProxyRequest
type Response = events.APIGatewayProxyResponse

type Handler interface {
	AddMetrics(ctx context.Context, req Request) (Response, error)
}

type handler struct {
	processor processor.Processor
}

func New(p processor.Processor) Handler {
	return &handler{
		processor: p,
	}
}

func (h *handler) AddMetrics(_ context.Context, req Request) (Response, error) {
	var r models.Request
	_ = json.Unmarshal([]byte(req.Body), &r)
	metric, err := h.processor.Process(r)
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
