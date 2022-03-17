package processor

import (
	"seb7887/go-trx-metrics/internal/repository"
	"seb7887/go-trx-metrics/pkg/models"
	"time"
)

type Processor interface {
	Process(req models.Request) (models.Metric, error)
}

type processor struct {
	repository repository.Repository
}

func New(r repository.Repository) Processor {
	return &processor{
		r,
	}
}

func (p *processor) Process(req models.Request) (models.Metric, error) {
	date := time.Now().Format(time.RFC3339)
	metric := models.Metric{
		AccountID: req.AccountID,
		Date:      date,
	}

	if err := p.repository.CreateMetric(metric); err != nil {
		return models.Metric{}, err
	}

	return metric, nil
}
