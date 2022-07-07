package model

import (
	"time"
)

type OperationHour struct {
	ID           uint   `gorm:"primaryKey"`
	RestaurantID uint   `gorm:"column:restaurant_id"`
	Day          string `gorm:"column:day"`
	OpenHour     string
	CloseHour    string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

// TableName overrides the table name
func (o OperationHour) TableName() string {
	return "ops_hour"
}
