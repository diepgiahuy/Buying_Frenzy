package storage

import (
	"context"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"gorm.io/gorm"
)

type Repo struct {
	Db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		Db: db,
	}
}

func (r *Repo) AddUser(ctx context.Context, user []model.User) error {
	if result := r.Db.WithContext(ctx).Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repo) AddUserWithBatches(ctx context.Context, user []model.User) error {

	if result := r.Db.WithContext(ctx).CreateInBatches(&user, 100); result.Error != nil {
		return result.Error
	}
	return nil
}
