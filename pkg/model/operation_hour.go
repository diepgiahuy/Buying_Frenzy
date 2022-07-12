package model

import (
	"context"
	"time"
)

type OperationHour struct {
	ID           *int64 `gorm:"primaryKey"`
	RestaurantID *int64 `gorm:"column:restaurant_id"`
	Day          string `gorm:"column:day"`
	OpenHour     string
	CloseHour    string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

type OperationHourStore interface {
	Add(ctx context.Context, restaurant []Restaurant) error
	DeleteByID(ctx context.Context, menu []Menu) error
}

// TableName overrides the table name
func (o OperationHour) TableName() string {
	return "ops_hour"
}
