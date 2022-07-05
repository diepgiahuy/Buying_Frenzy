package model

import (
	"context"
	"time"
)

type Menu struct {
	Data      Data `json:"menu"`
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Data []struct {
	DishName string  `json:"dishName"`
	Price    float64 `json:"price"`
}

type MenuStore interface {
	AddMenu(ctx context.Context, restaurant []Restaurant) error
}
