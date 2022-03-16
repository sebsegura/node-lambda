package models

type Request struct {
	AccountID string `json:"accountId" dynamodbav:"AccountId"`
	Amount    int    `json:"amount" dynamodbav:"Amount"`
}
