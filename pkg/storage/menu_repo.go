package storage

import (
	"errors"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type MenuStore struct {
	Db *gorm.DB
}

func NewMenuStore(db *gorm.DB) *MenuStore {
	return &MenuStore{
		Db: db,
	}
}
func (r *MenuStore) AddMenu(ctx context.Context, restaurant []model.Menu) error {
	if result := r.Db.WithContext(ctx).Table("menu").Create(&restaurant); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return result.Error
	}
	return nil
}
