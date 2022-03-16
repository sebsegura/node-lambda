package processor

import (
	"seb7887/go-trx-processor/internal/repository"
	"seb7887/go-trx-processor/pkg/models"
)

type Processor interface {
	Process(trx models.Request) error
}

type processor struct {
	repository repository.Repository
}

func New(r repository.Repository) Processor {
	return &processor{
		repository: r,
	}
}

func (p *processor) Process(trx models.Request) error {
	return p.repository.CreateTrx(trx)
}
