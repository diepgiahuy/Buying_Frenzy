package model

import (
	"context"
	"time"
)

type Restaurant struct {
	CashBalance    float64 `json:"cashBalance"`
	RestaurantName string  `json:"restaurantName" gorm:"column:name"`
	UpdatedAt      time.Time
	CreatedAt      time.Time
}

type RestaurantStore interface {
	AddRestaurant(ctx context.Context, restaurant []Restaurant) error
}
