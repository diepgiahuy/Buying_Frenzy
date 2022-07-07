package storage

import (
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"golang.org/x/net/context"
)

func (r *Repo) AddRestaurant(ctx context.Context, restaurant model.Restaurant) error {
	if result := r.Db.WithContext(ctx).Create(&restaurant); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repo) AddRestaurantWithBatches(ctx context.Context, restaurant []model.Restaurant) error {
	if result := r.Db.WithContext(ctx).CreateInBatches(&restaurant, 100); result.Error != nil {
		return result.Error
	}
	return nil
}
