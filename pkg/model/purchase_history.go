package model

import (
	"context"
	"time"
)

type PurchaseHistory struct {
	ID                uint    `gorm:"primaryKey"`
	UserId            *uint   `gorm:"column:user_id"`
	DishName          string  `json:"dishName"`
	RestaurantName    string  `json:"restaurantName"`
	TransactionAmount float64 `json:"transactionAmount"`
	TransactionDate   string  `json:"transactionDate"`
	UpdatedAt         time.Time
	CreatedAt         time.Time
}

// TableName overrides the table name
func (p PurchaseHistory) TableName() string {
	return "purchase_history"
}

type PurchaseHistoryStore interface {
	AddHistory(ctx context.Context, history PurchaseHistory) error
}
