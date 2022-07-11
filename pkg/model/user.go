package model

import (
	"context"
)

type User struct {
	ID              *int64            `json:"id,omitempty" gorm:"primaryKey"`
	CashBalance     float64           `json:"cashBalance"`
	Name            string            `json:"name"`
	PurchaseHistory []PurchaseHistory `json:"purchaseHistory"`
}

type UserStore interface {
	AddUserWithBatches(ctx context.Context, user []User) error
	GetUserByID(ctx context.Context, userID int64) (*User, error)
	DecreaseUserCashBalance(ctx context.Context, user *User, cash float64) error
}

// TableName overrides the table name
func (u User) TableName() string {
	return "user"
}
