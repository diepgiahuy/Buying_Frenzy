package model

import (
	"context"
	"time"
)

type OperationHour struct {
	CashBalance    float64 `json:"cashBalance"`
	OpeningHours   string  `json:"openingHours"`
	RestaurantName string  `json:"restaurantName"`
}
"id" SERIAL PRIMARY KEY,
"restaurant_id" integer not null,
"date" varchar NOT NULL,
"open_hour" timestamp NOT NULL ,
"close_hour" timestamp NOT NULL ,
"created_at" timestamptz NOT NULL DEFAULT (now()),
"updated_at" timestamptz NOT NULL DEFAULT (now()),

type Store interface {
	AddRestaurant(ctx context.Context, restaurant Restaurant)
	GetRestaurantByDateTime(ctx context.Context, DateTime time.Time)
}
