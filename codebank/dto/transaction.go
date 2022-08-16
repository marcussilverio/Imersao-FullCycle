package dto

import "time"
type Transaction struct {
	ID string
	Name string
	Number string
	ExpirationMonth int32
	ExpirationYear int32
	CVV int32
	Amount float64
	Store string
	CreatedAt time.Time
}