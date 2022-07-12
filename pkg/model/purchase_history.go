package model

import (
	"context"
	"time"
)

type PurchaseHistory struct {
	ID                int64   `gorm:"primaryKey"`
	UserId            *int64  `gorm:"column:user_id"`
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
	Add(ctx context.Context, history PurchaseHistory) (*PurchaseHistory, error)
	DeleteByID(ctx context.Context, historyId int64) error
}
