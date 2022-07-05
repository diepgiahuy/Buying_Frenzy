package storage

import (
	"errors"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"gorm.io/gorm"
)

type Repo struct {
	Db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *Repo {
	return &Repo{
		Db: db,
	}
}

func (r *Repo) InsertRestaurant(restaurant model.Restaurant) error {
	if result := r.Db.Table("restaurant").Save(&restaurant); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return result.Error
	}
	return nil
}
