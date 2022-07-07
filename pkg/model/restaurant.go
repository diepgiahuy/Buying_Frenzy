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
}

// TableName overrides the table name
func (r Restaurant) TableName() string {
	return "restaurant"
}
