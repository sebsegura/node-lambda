package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"seb7887/go-trx-processor/internal/processor"
	"seb7887/go-trx-processor/internal/repository"
	"seb7887/go-trx-processor/pkg/handler"
)

func main() {
	r := repository.New()
	p := processor.New(r)
	h := handler.New(p)

	lambda.Start(h.ProcessTrx)
}
