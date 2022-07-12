package storage

import (
	"context"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"github.com/diepgiahuy/Buying_Frenzy/util"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func createRandomOperationHour(t *testing.T) model.OperationHour {
	restaurant := CreateRandomRestaurantRecord(t)
	return model.OperationHour{
		ID:           util.RandomInt(20000, 21000),
		RestaurantID: restaurant.ID,
		Day:          DayParse[time.Now().Weekday().String()],
		OpenHour:     time.Now().Format("15:04:05"),
		CloseHour:    time.Now().Add(time.Duration(rand.Intn(1000))).Format("15:04:05"),
	}
}

func TestAddOperationHourRecordWithError(t *testing.T) {
	OperationHour := CreateRandomOperationHourRecord(t)
	_, err := operationHourStore.Add(context.Background(), OperationHour)
	require.Error(t, err)
	DeleteOperationHourRecord(t, *OperationHour.ID, *OperationHour.RestaurantID)
}

func CreateRandomOperationHourRecord(t *testing.T) model.OperationHour {
	OperationHour := createRandomOperationHour(t)
	got, err := operationHourStore.Add(context.Background(), OperationHour)
	require.NoError(t, err)
	require.NotEmpty(t, got)
	require.Equal(t, got.OpenHour, OperationHour.OpenHour)
	require.Equal(t, got.CloseHour, OperationHour.CloseHour)
	require.Equal(t, got.Day, OperationHour.Day)
	require.Equal(t, got.RestaurantID, OperationHour.RestaurantID)
	require.NotZero(t, got.ID)
	require.Equal(t, got.ID, OperationHour.ID)
	return OperationHour
}

func TestDeleteOperationHourById(t *testing.T) {
	OperationHour := createRandomOperationHour(t)
	err := operationHourStore.DeleteByID(context.Background(), *OperationHour.ID)
	require.NoError(t, err)
	DeleteOperationHourRecord(t, *OperationHour.ID, *OperationHour.RestaurantID)
}

func DeleteOperationHourRecord(t *testing.T, OperationHourId int64, restaurantId int64) {
	err := operationHourStore.DeleteByID(context.Background(), OperationHourId)
	require.NoError(t, err)
	DeleteRestaurantRecord(t, restaurantId)
}
