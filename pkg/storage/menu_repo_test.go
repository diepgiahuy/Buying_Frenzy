package storage

import (
	"context"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"github.com/diepgiahuy/Buying_Frenzy/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomMenu(t *testing.T) model.Menu {
	restaurant := CreateRandomRestaurantRecord(t)
	return model.Menu{
		ID:           util.RandomInt(20000, 21000),
		RestaurantID: restaurant.ID,
		DishName:     util.RandomOwner(),
		Price:        util.RandomMoney(),
	}
}

func TestAddMenuRecordWithError(t *testing.T) {
	menu := CreateRandomMenuRecord(t)
	_, err := menuStore.Add(context.Background(), menu)
	require.Error(t, err)
	DeleteMenuRecord(t, *menu.ID, *menu.RestaurantID)
}

func CreateRandomMenuRecord(t *testing.T) model.Menu {
	menu := createRandomMenu(t)
	got, err := menuStore.Add(context.Background(), menu)
	require.NoError(t, err)
	require.NotEmpty(t, got)
	require.Equal(t, got.Price, menu.Price)
	require.Equal(t, got.DishName, menu.DishName)
	require.Equal(t, got.RestaurantID, menu.RestaurantID)
	require.NotZero(t, got.ID)
	require.Equal(t, got.ID, menu.ID)
	return menu
}

func TestDeleteMenuById(t *testing.T) {
	menu := createRandomMenu(t)
	err := menuStore.DeleteByID(context.Background(), *menu.ID)
	require.NoError(t, err)
	DeleteMenuRecord(t, *menu.ID, *menu.RestaurantID)
}

func DeleteMenuRecord(t *testing.T, menuId int64, restaurantId int64) {
	err := menuStore.DeleteByID(context.Background(), menuId)
	require.NoError(t, err)
	DeleteRestaurantRecord(t, restaurantId)
}

func TestGetMenuByDishTerm(t *testing.T) {
	menu := CreateRandomMenuRecord(t)
	got, err := menuStore.GetByDishTerm(context.Background(), menu.DishName, 0, 1)
	require.NoError(t, err)
	require.NotEmpty(t, got)
	require.Equal(t, got[0].ID, menu.ID)
	DeleteMenuRecord(t, *menu.ID, *menu.RestaurantID)
}
func TestGetMenuByDishTermError(t *testing.T) {
	menu := CreateRandomMenuRecord(t)
	_, err := menuStore.GetByDishTerm(context.Background(), menu.DishName, -1, 1)
	require.Error(t, err)
	DeleteMenuRecord(t, *menu.ID, *menu.RestaurantID)

}
