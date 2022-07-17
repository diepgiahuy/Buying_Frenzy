package storage

import (
	"context"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"gorm.io/gorm"
	"log"
)

type HistoryStore struct {
	Db *gorm.DB
}

func NewHistoryStore(db *gorm.DB) *HistoryStore {
	return &HistoryStore{
		Db: db,
	}
}

// WithTx enables repository with transaction
func (h *HistoryStore) WithHistoryStoreTx(txHandle *gorm.DB) *HistoryStore {
	if txHandle == nil {
		log.Print("Transaction Database not found")
		return h
	}
	h.Db = txHandle
	return h
}

func (h *HistoryStore) Add(ctx context.Context, history model.PurchaseHistory) (*model.PurchaseHistory, error) {
	if result := h.Db.WithContext(ctx).Create(&history); result.Error != nil {
		return nil, result.Error
	}
	return &history, nil
}

func (h *HistoryStore) DeleteByID(ctx context.Context, historyId int64) error {
	result := h.Db.WithContext(ctx).Delete(&model.PurchaseHistory{}, historyId)
	return result.Error
}
