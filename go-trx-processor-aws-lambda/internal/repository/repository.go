package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"seb7887/go-trx-processor/internal/storage"
	"seb7887/go-trx-processor/pkg/models"
)

const TableName = "TestTrx"

type Repository interface {
	CreateTrx(trx models.Request) error
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

func (r *repository) CreateTrx(trx models.Request) error {
	item, err := dynamodbattribute.MarshalMap(trx)
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
