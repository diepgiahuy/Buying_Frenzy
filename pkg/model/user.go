package model

import (
	"context"
)

type User struct {
	ID              *uint             `json:"id,omitempty" gorm:"primaryKey" sql:"notnull"`
	CashBalance     float64           `json:"cashBalance"`
	Name            string            `json:"name"`
	PurchaseHistory []PurchaseHistory `json:"purchaseHistory"`
}

type UserStore interface {
	AddUser(ctx context.Context, user []User) error
	AddUserWithBatches(ctx context.Context, user []User) error
}

// TableName overrides the table name
func (u User) TableName() string {
	return "user"
}
