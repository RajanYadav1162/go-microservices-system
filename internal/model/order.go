package model

import (
	"time"

	"gorm.io/gorm"
)

type TicketOrder struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    string         `gorm:"not null"  json:"user_id"`
	ConcertID string         `gorm:"not null"  json:"concert_id"`
	Qty       int            `gorm:"not null"  json:"qty"`
	Amount    float64        `gorm:"not null"  json:"amount"`
	Status    string         `gorm:"not null"  json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (TicketOrder) TableName() string {
	return "ticket_orders"
}
