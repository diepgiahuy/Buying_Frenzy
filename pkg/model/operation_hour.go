package model

import (
	"time"
)

type opsHour struct {
	ID           uint `gorm:"primaryKey"`
	RestaurantID uint `gorm:"column:restaurant_id"`
	date         time.Weekday
	openHour     time.Time
	closeHour    time.Time
	Restaurant   Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DishName     string     `json:"dishName"`
	Price        float64    `json:"price"`
}

// TableName overrides the table name
func (o opsHour) TableName() string {
	return "ops_hour"
}
