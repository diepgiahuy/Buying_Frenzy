package storage

import (
	"errors"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func (r *Repo) AddRestaurant(ctx context.Context, restaurant []model.Restaurant) error {
	if result := r.Db.Table("restaurant").Create(&restaurant); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return result.Error
	}
	return nil
}
