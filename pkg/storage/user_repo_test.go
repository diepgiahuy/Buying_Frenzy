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

func creatRandomUser() model.User {
	return model.User{
		ID:          util.RandomInt(8000, 20000),
		CashBalance: util.RandomMoney(),
		Name:        util.RandomOwner(),
	}
}

func CreateUserRecord(t *testing.T) model.User {
	user := creatRandomUser()
	arg, err := userStore.AddUser(context.Background(), user)
	require.NoError(t, err)
	require.NotEmpty(t, arg)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.CashBalance, user.CashBalance)
	require.NotZero(t, user.ID)
	require.Equal(t, arg.ID, user.ID)
	return user
}

func DeleteUserRecord(t *testing.T, userID int64) {
	err := userStore.DeleteUserByID(context.Background(), userID)
	require.NoError(t, err)
}

func TestAddUserRecordWithError(t *testing.T) {
	user := CreateUserRecord(t)
	_, err := userStore.AddUser(context.Background(), user)
	require.Error(t, err)
	DeleteUserRecord(t, *user.ID)
}

func TestDeleteUserById(t *testing.T) {
	account1 := CreateUserRecord(t)
	err := userStore.DeleteUserByID(context.Background(), *account1.ID)
	require.NoError(t, err)
	account2, err := userStore.GetByID(context.Background(), *account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	require.Empty(t, account2)
}

func TestGetUserById(t *testing.T) {
	user1 := CreateUserRecord(t)
	user2 := creatRandomUser()
	tests := []struct {
		name    string
		user    model.User
		wantErr error
	}{
		{"querySuccess", user1, nil},
		{"recordNotFound", user2, gorm.ErrRecordNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := userStore.GetByID(context.Background(), *tt.user.ID)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
				return
			}
			require.NoError(t, err)
			require.NotEmpty(t, got)
			require.Equal(t, tt.user.ID, got.ID)
			require.Equal(t, tt.user.CashBalance, got.CashBalance)
			require.Equal(t, tt.user.Name, got.Name)
			DeleteUserRecord(t, *user1.ID)
		})
	}
}

func TestAddUserWithBatches(t *testing.T) {

	var users []model.User
	user1 := creatRandomUser()
	users = append(users, user1)
	tests := []struct {
		name    string
		user    model.User
		wantErr error
	}{
		{"querySuccess", user1, nil},
		{"duplicateError", user1, errors.New("ERROR: duplicate key value violates unique constraint \"user_pkey\" (SQLSTATE 23505)")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userStore.AddWithBatches(context.Background(), users)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.wantErr.Error())
				return
			}
			require.NoError(t, err)
		})
	}
	err := userStore.DeleteUserByID(context.Background(), *user1.ID)
	require.NoError(t, err)
}

func TestDecreaseUserCashBalance(t *testing.T) {
	user1 := CreateUserRecord(t)
	cash := 100.0
	tests := []struct {
		name    string
		user    model.User
		wantErr error
	}{
		{"querySuccess", user1, nil},
		{"foundError", creatRandomUser(), gorm.ErrMissingWhereClause},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr != nil {
				err := userStore.DecreaseCashBalance(context.Background(), nil, cash)
				require.Error(t, err)
				require.EqualError(t, err, tt.wantErr.Error())
				return
			}
			err := userStore.DecreaseCashBalance(context.Background(), &tt.user, cash)
			require.NoError(t, err)
			got, err := userStore.GetByID(context.Background(), *tt.user.ID)
			require.NoError(t, err)
			require.NotEmpty(t, got)
			require.Equal(t, tt.user.CashBalance-cash, got.CashBalance)
		})
	}
	err := userStore.DeleteUserByID(context.Background(), *user1.ID)
	require.NoError(t, err)

}
