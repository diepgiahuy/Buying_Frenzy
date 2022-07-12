package storage

import (
	"fmt"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
)

type RestaurantStore struct {
	Db *gorm.DB
}

func NewRestaurantStore(db *gorm.DB) *RestaurantStore {
	return &RestaurantStore{
		Db: db,
	}
}

func (r *RestaurantStore) Add(ctx context.Context, restaurant model.Restaurant) (*model.Restaurant, error) {
	if result := r.Db.WithContext(ctx).Create(&restaurant); result.Error != nil {
		return nil, result.Error
	}
	return &restaurant, nil
}

func (r *RestaurantStore) AddWithBatches(ctx context.Context, restaurant []model.Restaurant) error {
	if result := r.Db.WithContext(ctx).CreateInBatches(&restaurant, 100); result.Error != nil {
		return result.Error
	}
	return nil
}

var DayParse = map[string]string{
	"Monday":    "Mon",
	"Tuesday":   "Tues",
	"Wednesday": "Weds",
	"Thursday":  "Thurs",
	"Friday":    "Fri",
	"Saturday":  "Sat",
	"Sunday":    "Sun",
}

func (r *RestaurantStore) GetByDate(ctx context.Context, date string, offset int, pageSize int) ([]model.Restaurant, error) {
	parse, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return nil, err
	}
	day := DayParse[parse.Weekday().String()]
	times := fmt.Sprintf("%02d:%02d:00", parse.Hour(), parse.Minute())
	var res []model.Restaurant
	if result := r.Db.WithContext(ctx).Preload("OperationHour").Preload("Menu").
		Joins("Inner join ops_hour on ops_hour.restaurant_id = restaurant.id").
		Joins("Inner join menu on menu.restaurant_id = restaurant.id").
		Where("open_hour = ? and day = ?", times, day).Group("restaurant.id").Offset(offset).Limit(pageSize).Find(&res); result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func (r *RestaurantStore) GetByCompareGreaterDish(ctx context.Context, lowPrice float64, highPrice float64, numDishes int, topList int) ([]model.Restaurant, error) {

	var res []model.Restaurant
	if result := r.Db.WithContext(ctx).Raw("SELECT * FROM  restaurant"+
		"\nwhere id IN ("+
		"\n    Select restaurant_id"+
		"\n    from ("+
		"\n             SELECT restaurant_id, count(*) OVER (PARTITION BY restaurant_id) as num_dishes"+
		"\n             from menu"+
		"\n             where price between ? and ?"+
		"\n         ) as restaurant"+
		"\n    where restaurant.num_dishes >  ?"+
		"\n    group by restaurant_id,num_dishes"+
		"\n    order by num_dishes desc"+
		"\n    limit ?"+
		"\n)"+
		"\n order by  name asc", lowPrice, highPrice, numDishes, topList).Preload("OperationHour").
		Preload("Menu").
		Find(&res); result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func (r *RestaurantStore) GetByCompareLesserDish(ctx context.Context, lowPrice float64, highPrice float64, numDishes int, topList int) ([]model.Restaurant, error) {

	var res []model.Restaurant
	if result := r.Db.WithContext(ctx).Raw("SELECT * FROM  restaurant"+
		"\nwhere id IN ("+
		"\n    Select restaurant_id"+
		"\n    from ("+
		"\n             SELECT restaurant_id, count(*) OVER (PARTITION BY restaurant_id) as num_dishes"+
		"\n             from menu"+
		"\n             where price between ? and ?"+
		"\n         ) as restaurant"+
		"\n    where restaurant.num_dishes < ?"+
		"\n    group by restaurant_id,num_dishes"+
		"\n    order by num_dishes desc"+
		"\n    limit ?"+
		"\n)"+
		"\n order by  name asc", lowPrice, highPrice, numDishes, topList).
		Preload("OperationHour").
		Preload("Menu").
		Find(&res); result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func (r *RestaurantStore) GetByTerm(ctx context.Context, term string, offset int, pageSize int) ([]model.Restaurant, error) {

	var res []model.Restaurant
	if result := r.Db.WithContext(ctx).Raw(
		"SELECT * FROM ("+
			"\n         SELECT DISTINCT name,id,cash_balance"+
			"\n         FROM restaurant"+
			"\n         WHERE name ILIKE ?"+
			"\n     ) alias"+
			"\n     ORDER BY name ILIKE ? DESC, name"+
			"\n 	OFFSET ?"+
			"\n 	LIMIT  ?", "%"+term+"%", term+"%", offset, pageSize).Preload("OperationHour").Preload("Menu").
		Joins("Inner join ops_hour on ops_hour.restaurant_id = restaurant.id").
		Joins("Inner join menu on menu.restaurant_id = restaurant.id").Find(&res); result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func (r *RestaurantStore) GetByID(ctx context.Context, restaurantID int64) (*model.Restaurant, error) {

	var restaurantData *model.Restaurant
	if result := r.Db.WithContext(ctx).Preload("Menu").First(&restaurantData, restaurantID); result.Error != nil {
		return nil, result.Error
	}
	return restaurantData, nil
}

func (r *RestaurantStore) IncreaseCashBalance(ctx context.Context, restaurant *model.Restaurant, cash float64) error {

	if result := r.Db.WithContext(ctx).Model(&restaurant).Update("cash_balance", gorm.Expr("cash_balance + ?", cash)); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RestaurantStore) DeleteByID(ctx context.Context, restaurant int64) error {
	result := r.Db.WithContext(ctx).Delete(&model.Restaurant{}, restaurant)
	return result.Error
}
