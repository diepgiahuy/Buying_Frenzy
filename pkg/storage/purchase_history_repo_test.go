package storage

import (
	"context"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"github.com/diepgiahuy/Buying_Frenzy/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomHistory(t *testing.T) (model.PurchaseHistory, int64) {
	restaurant := CreateRandomRestaurantRecord(t)
	user := CreateUserRecord(t)
	return model.PurchaseHistory{
		ID:                *util.RandomInt(15000, 20000),
		UserId:            user.ID,
		DishName:          util.RandomOwner(),
		RestaurantName:    restaurant.RestaurantName,
		TransactionAmount: util.RandomMoney(),
		TransactionDate:   time.Now().Format(time.RFC3339),
	}, *restaurant.ID
}

func TestAddHistoryRecordWithError(t *testing.T) {
	history, restaurantId := CreateRandomHistoryRecord(t)
	_, err := historyStore.Add(context.Background(), history)
	require.Error(t, err)
	DeleteHistoryRecord(t, history.ID, *history.UserId, restaurantId)
}

func CreateRandomHistoryRecord(t *testing.T) (model.PurchaseHistory, int64) {
	history, restaurantId := createRandomHistory(t)
	got, err := historyStore.Add(context.Background(), history)
	require.NoError(t, err)
	require.NotEmpty(t, got)
	require.Equal(t, got.RestaurantName, history.RestaurantName)
	require.Equal(t, got.DishName, history.DishName)
	require.Equal(t, got.TransactionAmount, history.TransactionAmount)
	require.Equal(t, got.TransactionDate, history.TransactionDate)
	require.Equal(t, got.UserId, history.UserId)
	require.NotZero(t, got.ID)
	return history, restaurantId
}

func TestDeleteHistoryById(t *testing.T) {
	historyRecord, restaurantID := createRandomHistory(t)
	err := historyStore.DeleteByID(context.Background(), historyRecord.ID)
	require.NoError(t, err)
	DeleteHistoryRecord(t, historyRecord.ID, *historyRecord.UserId, restaurantID)
}

func DeleteHistoryRecord(t *testing.T, historyID int64, userId int64, restaurantId int64) {
	err := historyStore.DeleteByID(context.Background(), historyID)
	require.NoError(t, err)
	DeleteRestaurantRecord(t, restaurantId)
	DeleteUserRecord(t, userId)
}
