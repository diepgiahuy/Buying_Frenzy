package storage

import (
	"context"
	"errors"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"gorm.io/gorm"
)

type UserStore struct {
	Db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		Db: db,
	}
}

func (r *UserStore) AddUser(ctx context.Context, user []model.User) error {
	if result := r.Db.WithContext(ctx).Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserStore) AddUserWithBatches(ctx context.Context, user []model.User) error {

	if result := r.Db.WithContext(ctx).CreateInBatches(&user, 100); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserStore) GetUserByID(ctx context.Context, userID int64) (*model.User, error) {
	var userData *model.User
	if result := r.Db.WithContext(ctx).First(&userData, userID); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}
	return userData, nil
}

func (r *UserStore) DecreaseUserCashBalance(ctx context.Context, user *model.User, cash float64) error {
	if result := r.Db.WithContext(ctx).Model(&user).Update("cash_balance", gorm.Expr("cash_balance - ?", cash)); result.Error != nil {
		return result.Error
	}
	return nil
}
