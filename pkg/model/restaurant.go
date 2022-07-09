package model

import (
	"context"
	"time"
)

type Restaurant struct {
	ID             uint    `gorm:"primaryKey"`
	CashBalance    float64 `json:"cashBalance"`
	RestaurantName string  `json:"restaurantName" gorm:"column:name"`
	Menu           []Menu  `json:"menu" gorm:"foreignKey:RestaurantID"`
	OperationHour  []OperationHour
	UpdatedAt      time.Time
	CreatedAt      time.Time
}

type RestaurantStore interface {
	AddRestaurant(ctx context.Context, restaurant Restaurant) error
	AddRestaurantWithBatches(ctx context.Context, restaurant []Restaurant) error
	GetRestaurantWithDate(ctx context.Context, date string, offset int, pageSize int) ([]Restaurant, error)
	GetRestaurantWithCompareMore(ctx context.Context, priceBot float32, priceTop float32, numDishes int, top int) ([]Restaurant, error)
	GetRestaurantWithCompareLess(ctx context.Context, priceBot float32, priceTop float32, numDishes int, top int) ([]Restaurant, error)
	GetRestaurantByTerm(ctx context.Context, term string, offset int, pageSize int) ([]Restaurant, error)
	GetRestaurantByID(ctx context.Context, restaurantID int64) (*Restaurant, error)
}

// TableName overrides the table name
func (r Restaurant) TableName() string {
	return "restaurant"
}
