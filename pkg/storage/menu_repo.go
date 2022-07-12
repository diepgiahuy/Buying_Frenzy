package storage

import (
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
func (r *MenuStore) Add(ctx context.Context, menu model.Menu) (*model.Menu, error) {
	if result := r.Db.WithContext(ctx).Table("menu").Create(&menu); result.Error != nil {
		return nil, result.Error
	}
	return &menu, nil
}

func (r *MenuStore) DeleteByID(ctx context.Context, menuID int64) error {
	result := r.Db.WithContext(ctx).Delete(&model.Menu{}, menuID)
	return result.Error
}

func (r *MenuStore) GetByDishTerm(ctx context.Context, term string, offset int, pageSize int) ([]model.Menu, error) {

	var res []model.Menu
	if result := r.Db.WithContext(ctx).Raw("SELECT * from menu "+
		"\n INNER JOIN ("+
		"\n    SELECT * FROM ("+
		"\n             SELECT DISTINCT dish_name, id"+
		"\n             FROM menu"+
		"\n             WHERE dish_name ILIKE ?"+
		"\n         ) alias"+
		"\n    ORDER BY dish_name ILIKE ? DESC, dish_name"+
		"\n) as Alias on Alias.id = menu.id"+
		"\n 	OFFSET ?"+
		"\n 	LIMIT  ?", "%"+term+"%", term+"%", offset, pageSize).Find(&res); result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}
