package dto

import "time"

type Transaction struct {
	ID              string
	Name            string
	Number          string
	ExpirationMonth string
	ExpirationYear  string
	CVV             string
	Amount          float64
	Store           string
	Description     string
	CreatedAt       time.Time
}
