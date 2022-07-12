package model

import (
	"context"
)

type Menu struct {
	ID           *int64  `gorm:"primaryKey"`
	RestaurantID *int64  `gorm:"column:restaurant_id"`
	DishName     string  `json:"dishName"`
	Price        float64 `json:"price"`
}

type MenuStore interface {
	Add(ctx context.Context, menu Menu) (*Menu, error)
	DeleteByID(ctx context.Context, menuID int64) error
	GetByDishTerm(ctx context.Context, term string, offset int, pageSize int) ([]Menu, error)
}

// TableName overrides the table name
func (m Menu) TableName() string {
	return "menu"
}
