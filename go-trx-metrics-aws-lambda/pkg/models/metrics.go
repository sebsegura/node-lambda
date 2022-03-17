package models

import "encoding/json"

type Metric struct {
	AccountID string `json:"accountId" dynamodbav:"AccountId"`
	Date      string `json:"date" dynamodbav:"Date"`
}

func (m *Metric) ToJsonStr() string {
	b, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(b)
}
