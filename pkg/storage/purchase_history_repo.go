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

func (r *HistoryStore) AddHistory(ctx context.Context, history model.PurchaseHistory) (*model.PurchaseHistory, error) {
	if result := r.Db.WithContext(ctx).Create(&history); result.Error != nil {
		return nil, result.Error
	}
	return &history, nil
}

func (r *HistoryStore) DeleteHistoryByID(ctx context.Context, historyId int64) error {
	result := r.Db.WithContext(ctx).Delete(&model.PurchaseHistory{}, historyId)
	return result.Error
}
