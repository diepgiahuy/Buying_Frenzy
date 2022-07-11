package storage

import (
	"context"
	"errors"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"github.com/diepgiahuy/Buying_Frenzy/util"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
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
	got, err := restaurantStore.AddRestaurant(context.Background(), restaurant)
	require.NoError(t, err)
	require.NotEmpty(t, got)

	require.Equal(t, got.RestaurantName, restaurant.RestaurantName)
	require.Equal(t, got.CashBalance, restaurant.CashBalance)
	require.NotZero(t, got.ID)
	require.Equal(t, got.ID, restaurant.ID)
	return restaurant
}

func TestRestaurantWithBatches(t *testing.T) {

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
			err := restaurantStore.AddRestaurantWithBatches(context.Background(), tt.restaurant)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.wantErr.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestAddRestaurantRecordWithError(t *testing.T) {
	restaurant := CreateRandomRestaurantRecord(t)
	_, err := restaurantStore.AddRestaurant(context.Background(), restaurant)
	require.Error(t, err)
	DeleteRestaurantRecord(t, *restaurant.ID)
}

func TestDeleteRestaurantById(t *testing.T) {
	restaurantRecord := CreateRandomRestaurantRecord(t)
	err := restaurantStore.DeleteRestaurantByID(context.Background(), *restaurantRecord.ID)
	require.NoError(t, err)
	got, err := restaurantStore.GetRestaurantByID(context.Background(), *restaurantRecord.ID)
	require.Error(t, err)
	require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	require.Empty(t, got)
}

func DeleteRestaurantRecord(t *testing.T, restaurantId int64) {
	err := restaurantStore.DeleteRestaurantByID(context.Background(), restaurantId)
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
				err := restaurantStore.IncreaseRestaurantCashBalance(context.Background(), nil, cash)
				require.Error(t, err)
				require.EqualError(t, err, tt.wantErr.Error())
				return
			}
			err := restaurantStore.IncreaseRestaurantCashBalance(context.Background(), &tt.restaurant, cash)
			require.NoError(t, err)
			got, err := restaurantStore.GetRestaurantByID(context.Background(), *tt.restaurant.ID)
			require.NoError(t, err)
			require.NotEmpty(t, got)
			require.Equal(t, tt.restaurant.CashBalance+cash, got.CashBalance)
		})
	}
	err := restaurantStore.DeleteRestaurantByID(context.Background(), *restaurant.ID)
	require.NoError(t, err)

}
