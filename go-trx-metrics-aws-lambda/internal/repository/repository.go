package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"seb7887/go-trx-metrics/internal/storage"
	"seb7887/go-trx-metrics/pkg/models"
)

const TableName = "TestMetrics"

type Repository interface {
	CreateMetric(m models.Metric) error
}

type repository struct {
	db    *dynamodb.DynamoDB
	table string
}

func New() Repository {
	db := storage.NewDynamoDbClient()
	return &repository{
		db:    db,
		table: TableName,
	}
}

func (r *repository) CreateMetric(m models.Metric) error {
	item, err := dynamodbattribute.MarshalMap(m)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(r.table),
	}

	_, err = r.db.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
