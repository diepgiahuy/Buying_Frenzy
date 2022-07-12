package storage

import (
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type OperationHourStore struct {
	Db *gorm.DB
}

func NewOperationHourStore(db *gorm.DB) *OperationHourStore {
	return &OperationHourStore{
		Db: db,
	}
}
func (r *OperationHourStore) Add(ctx context.Context, opsHour model.OperationHour) (*model.OperationHour, error) {
	if result := r.Db.WithContext(ctx).Create(&opsHour); result.Error != nil {
		return nil, result.Error
	}
	return &opsHour, nil
}

func (r *OperationHourStore) DeleteByID(ctx context.Context, opsHourId int64) error {
	result := r.Db.WithContext(ctx).Delete(&model.OperationHour{}, opsHourId)
	return result.Error
}
