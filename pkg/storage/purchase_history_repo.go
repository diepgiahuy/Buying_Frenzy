package storage

import (
	"context"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
)

func (r *Repo) AddHistory(ctx context.Context, history model.PurchaseHistory) error {
	if result := r.Db.WithContext(ctx).Create(&history); result.Error != nil {
		return result.Error
	}
	return nil
}
