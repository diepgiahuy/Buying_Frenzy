package storage

import (
	"context"
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

func (r *UserStore) AddUser(ctx context.Context, user model.User) (*model.User, error) {
	if result := r.Db.WithContext(ctx).Create(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserStore) AddWithBatches(ctx context.Context, user []model.User) error {
	if result := r.Db.WithContext(ctx).CreateInBatches(&user, 100); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserStore) GetByID(ctx context.Context, userID int64) (*model.User, error) {
	var userData *model.User
	if result := r.Db.WithContext(ctx).First(&userData, userID); result.Error != nil {
		return nil, result.Error
	}
	return userData, nil
}

func (r *UserStore) DeleteUserByID(ctx context.Context, userID int64) error {
	result := r.Db.WithContext(ctx).Delete(&model.User{}, userID)
	return result.Error
}

func (r *UserStore) DecreaseCashBalance(ctx context.Context, user *model.User, cash float64) error {
	if result := r.Db.WithContext(ctx).Model(&user).Update("cash_balance", gorm.Expr("cash_balance - ?", cash)); result.Error != nil {
		return result.Error
	}
	return nil
}
