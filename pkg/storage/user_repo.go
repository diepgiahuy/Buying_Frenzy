package storage

import (
	"context"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"gorm.io/gorm"
	"log"
)

type UserStore struct {
	Db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		Db: db,
	}
}

// WithUserTx enables repository with transaction
func (u *UserStore) WithUserTx(txHandle *gorm.DB) *UserStore {
	if txHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.Db = txHandle
	return u
}

func (u *UserStore) AddUser(ctx context.Context, user model.User) (*model.User, error) {
	if result := u.Db.WithContext(ctx).Create(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *UserStore) AddWithBatches(ctx context.Context, user []model.User) error {
	if result := u.Db.WithContext(ctx).CreateInBatches(&user, 100); result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserStore) GetByID(ctx context.Context, userID int64) (*model.User, error) {
	var userData *model.User
	if result := u.Db.WithContext(ctx).First(&userData, userID); result.Error != nil {
		return nil, result.Error
	}
	return userData, nil
}

func (u *UserStore) DeleteUserByID(ctx context.Context, userID int64) error {
	result := u.Db.WithContext(ctx).Delete(&model.User{}, userID)
	return result.Error
}

func (u *UserStore) DecreaseCashBalance(ctx context.Context, user *model.User, cash float64) error {
	if result := u.Db.WithContext(ctx).Model(&user).Update("cash_balance", gorm.Expr("cash_balance - ?", cash)); result.Error != nil {
		return result.Error
	}
	return nil
}
