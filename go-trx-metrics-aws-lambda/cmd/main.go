package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"seb7887/go-trx-metrics/internal/processor"
	"seb7887/go-trx-metrics/internal/repository"
	"seb7887/go-trx-metrics/pkg/handler"
)

func main() {
	r := repository.New()
	p := processor.New(r)
	h := handler.New(p)

	lambda.Start(h.AddMetrics)
}
