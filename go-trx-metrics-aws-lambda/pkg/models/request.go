package models

type Request struct {
	AccountID string `json:"accountId"`
	Amount    int    `json:"amount"`
}
