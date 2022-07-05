package restaurant

import (
	"context"
	"time"
)

type Model struct {
	CashBalance float64 `json:"cashBalance"`
	Menu        []struct {
		DishName string  `json:"dishName"`
		Price    float64 `json:"price"`
	} `json:"menu"`
	OpeningHours   string `json:"openingHours"`
	RestaurantName string `json:"restaurantName"`
}

type Store interface {
	AddRestaurant(ctx context.Context, model Model)
	GetRestaurantByDateTime(ctx context.Context, DateTime time.Time)
}

func TransformData() {

}