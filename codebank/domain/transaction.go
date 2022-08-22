package domain

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCardId string
	CreatedAt    time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()
	return t
}
