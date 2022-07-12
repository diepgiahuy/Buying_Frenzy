package model

import (
	"context"
	"time"
)

type Restaurant struct {
	ID             *int64  `gorm:"primaryKey"`
	CashBalance    float64 `json:"cashBalance"`
	RestaurantName string  `json:"restaurantName" gorm:"column:name"`
	Menu           []Menu  `json:"menu" gorm:"foreignKey:RestaurantID"`
	OperationHour  []OperationHour
	UpdatedAt      time.Time
	CreatedAt      time.Time
}

type RestaurantStore interface {
	Add(ctx context.Context, restaurant Restaurant) (*Restaurant, error)
	AddWithBatches(ctx context.Context, restaurant []Restaurant) error
	GetByDate(ctx context.Context, date string, offset int, pageSize int) ([]Restaurant, error)
	GetByCompareGreaterDish(ctx context.Context, priceBot float64, priceTop float64, numDishes int, top int) ([]Restaurant, error)
	GetByCompareLesserDish(ctx context.Context, priceBot float64, priceTop float64, numDishes int, top int) ([]Restaurant, error)
	GetByTerm(ctx context.Context, term string, offset int, pageSize int) ([]Restaurant, error)
	GetByID(ctx context.Context, restaurantID int64) (*Restaurant, error)
	IncreaseCashBalance(ctx context.Context, restaurant *Restaurant, cash float64) error
	DeleteByID(ctx context.Context, restaurant int64) error
}

// TableName overrides the table name
func (r Restaurant) TableName() string {
	return "restaurant"
}
