package model

import (
	"context"
)

type Menu struct {
	ID           uint    `gorm:"primaryKey"`
	RestaurantID uint    `gorm:"column:restaurant_id"`
	DishName     string  `json:"dishName"`
	Price        float64 `json:"price"`
}

type MenuStore interface {
	AddMenu(ctx context.Context, restaurant []Restaurant) error
}

// TableName overrides the table name
func (m Menu) TableName() string {
	return "menu"
}
