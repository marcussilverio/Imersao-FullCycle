package dto

import "time"
type Transaction struct {
	ID string `json:"transaction_id"`
	Name string `json:"-"`
	Number string `json:"credit_card_number"`
	ExpirationMonth int32 `json:"-"`
	ExpirationYear int32 `json:"-"`
	CVV int32 `json:"-"`
	Amount float64 `json:"Amount"`
	Store string `json:"Store"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"payment_date"`
}