package storage

import (
	"context"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"gorm.io/gorm"
)

type HistoryStore struct {
	Db *gorm.DB
}

func NewHistoryStore(db *gorm.DB) *HistoryStore {
	return &HistoryStore{
		Db: db,
	}
}

func (r *HistoryStore) AddHistory(ctx context.Context, history model.PurchaseHistory) error {
	if result := r.Db.WithContext(ctx).Create(&history); result.Error != nil {
		return result.Error
	}
	return nil
}
