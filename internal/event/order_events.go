package event

import "time"

type OrderCreated struct {
	OrderID   string    `json:"order_id"`
	UserID    string    `json:"user_id"`
	ConcertID string    `json:"concert_id"`
	Qty       int       `json:"qty"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
