package storage

import (
	"context"
	"errors"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"github.com/diepgiahuy/Buying_Frenzy/util"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"math/rand"
	"testing"
	"time"
)

func createRandomRestaurant() model.Restaurant {
	return model.Restaurant{
		ID:             util.RandomInt(3000, 5000),
		CashBalance:    util.RandomMoney(),
		RestaurantName: util.RandomOwner(),
	}
}

func CreateRandomRestaurantRecord(t *testing.T) model.Restaurant {
	restaurant := createRandomRestaurant()
	got, err := restaurantStore.Add(context.Background(), restaurant)
	require.NoError(t, err)
	require.NotEmpty(t, got)

	require.Equal(t, got.RestaurantName, restaurant.RestaurantName)
	require.Equal(t, got.CashBalance, restaurant.CashBalance)
	require.NotZero(t, got.ID)
	require.Equal(t, got.ID, restaurant.ID)
	return restaurant
}

func TestAddRestaurantWithBatches(t *testing.T) {
	var restaurants []model.Restaurant
	restaurant1 := createRandomRestaurant()
	restaurants = append(restaurants, restaurant1)
	tests := []struct {
		name       string
		restaurant []model.Restaurant
		wantErr    error
	}{
		{"querySuccess", restaurants, nil},
		{"duplicateError", restaurants, errors.New("ERROR: duplicate key value violates unique constraint \"restaurant_pkey\" (SQLSTATE 23505)")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := restaurantStore.AddWithBatches(context.Background(), tt.restaurant)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.wantErr.Error())
				return
			}
			require.NoError(t, err)
		})
	}
	DeleteRestaurantRecord(t, *restaurant1.ID)

}

func TestAddRestaurantRecordWithError(t *testing.T) {
	restaurant := CreateRandomRestaurantRecord(t)
	_, err := restaurantStore.Add(context.Background(), restaurant)
	require.Error(t, err)
	DeleteRestaurantRecord(t, *restaurant.ID)
}

func TestDeleteRestaurantById(t *testing.T) {
	restaurantRecord := CreateRandomRestaurantRecord(t)
	err := restaurantStore.DeleteByID(context.Background(), *restaurantRecord.ID)
	require.NoError(t, err)
	got, err := restaurantStore.GetByID(context.Background(), *restaurantRecord.ID)
	require.Error(t, err)
	require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	require.Empty(t, got)
}

func DeleteRestaurantRecord(t *testing.T, restaurantId int64) {
	err := restaurantStore.DeleteByID(context.Background(), restaurantId)
	require.NoError(t, err)
}

func TestIncreaseRestaurantCashBalance(t *testing.T) {
	restaurant := CreateRandomRestaurantRecord(t)
	cash := 100.0
	tests := []struct {
		name       string
		restaurant model.Restaurant
		wantErr    error
	}{
		{"querySuccess", restaurant, nil},
		{"foundError", createRandomRestaurant(), gorm.ErrMissingWhereClause},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr != nil {
				err := restaurantStore.IncreaseCashBalance(context.Background(), nil, cash)
				require.Error(t, err)
				require.EqualError(t, err, tt.wantErr.Error())
				return
			}
			err := restaurantStore.IncreaseCashBalance(context.Background(), &tt.restaurant, cash)
			require.NoError(t, err)
			got, err := restaurantStore.GetByID(context.Background(), *tt.restaurant.ID)
			require.NoError(t, err)
			require.NotEmpty(t, got)
			require.Equal(t, tt.restaurant.CashBalance+cash, got.CashBalance)
		})
	}
	err := restaurantStore.DeleteByID(context.Background(), *restaurant.ID)
	require.NoError(t, err)
}

func TestGetRestaurantByTerm(t *testing.T) {
	restaurant := CreateRandomRestaurantRecord(t)
	got, err := restaurantStore.GetByTerm(context.Background(), restaurant.RestaurantName, 0, 1)
	require.NoError(t, err)
	require.NotEmpty(t, got)
	require.Equal(t, got[0].ID, restaurant.ID)
	DeleteRestaurantRecord(t, *restaurant.ID)
}

func TestGetRestaurantByTermError(t *testing.T) {
	restaurant := CreateRandomRestaurantRecord(t)
	_, err := restaurantStore.GetByTerm(context.Background(), restaurant.RestaurantName, -1, 1)
	require.Error(t, err)
	DeleteRestaurantRecord(t, *restaurant.ID)
}

func TestGetRestaurantByDate(t *testing.T) {
	restaurant := model.Restaurant{
		ID:             util.RandomInt(3000, 5000),
		CashBalance:    util.RandomMoney(),
		RestaurantName: util.RandomOwner(),
	}
	date := time.Now()
	opsHour := model.OperationHour{
		ID:           util.RandomInt(20000, 21000),
		RestaurantID: restaurant.ID,
		Day:          DayParse[date.Weekday().String()],
		OpenHour:     date.Format("15:04:00"),
		CloseHour:    date.Add(time.Duration(rand.Intn(1000))).Format("15:04:00"),
	}
	menu := model.Menu{
		ID:           util.RandomInt(20000, 21000),
		RestaurantID: restaurant.ID,
		DishName:     util.RandomOwner(),
		Price:        util.RandomMoney(),
	}
	restaurant.Menu = []model.Menu{menu}
	restaurant.OperationHour = []model.OperationHour{opsHour}
	got, err := restaurantStore.Add(context.Background(), restaurant)
	require.NoError(t, err)
	require.NotEmpty(t, got)
	require.Equal(t, got.RestaurantName, restaurant.RestaurantName)
	require.Equal(t, got.CashBalance, restaurant.CashBalance)
	require.Equal(t, got.OperationHour[0].ID, opsHour.ID)
	require.Equal(t, got.Menu[0].ID, menu.ID)
	require.NotZero(t, got.ID)
	require.Equal(t, got.ID, restaurant.ID)
	restaurants, err := restaurantStore.GetByDate(context.Background(), date.Format("2006-01-02 15:04:05"), 0, 10)
	require.NoError(t, err)
	require.NotEmpty(t, restaurants)
	isEqual := false
	for _, res := range restaurants {
		if *res.ID == *restaurant.ID {
			isEqual = true
		}
	}
	require.Equal(t, true, isEqual)
	require.Equal(t, restaurants[0].ID, restaurant.ID)
	err = menuStore.DeleteByID(context.Background(), *menu.ID)
	DeleteOperationHourRecord(t, *opsHour.ID, *restaurant.ID)
	require.NoError(t, err)
}

func TestGetRestaurantByDateWithError_DateFormat(t *testing.T) {
	_, err := restaurantStore.GetByDate(context.Background(), "", 0, 10)
	require.Error(t, err)
}

func TestGetByCompareLesserDish(t *testing.T) {
	menu := CreateRandomMenuRecord(t)
	got, err := restaurantStore.GetByCompareLesserDish(context.Background(), menu.Price-10, menu.Price+10, 2, 100)
	require.NoError(t, err)
	require.NotEmpty(t, got)
	isEqual := false
	for _, restaurant := range got {
		if *menu.RestaurantID == *restaurant.ID {
			isEqual = true
		}
	}
	require.Equal(t, true, isEqual)
	DeleteMenuRecord(t, *menu.ID, *menu.RestaurantID)
}

func TestGetByCompareGreaterDish(t *testing.T) {
	menu := CreateRandomMenuRecord(t)
	got, err := restaurantStore.GetByCompareGreaterDish(context.Background(), menu.Price-10, menu.Price+10, 0, 100)
	require.NoError(t, err)
	require.NotEmpty(t, got)
	isEqual := false
	for _, restaurant := range got {
		if *menu.RestaurantID == *restaurant.ID {
			isEqual = true
		}
	}
	require.Equal(t, true, isEqual)
	DeleteMenuRecord(t, *menu.ID, *menu.RestaurantID)
}
