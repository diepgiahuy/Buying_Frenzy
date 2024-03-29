package api

import (
	"fmt"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type purchaseRequest struct {
	UserID       *int64 `json:"user_id" binding:"required,min=0,numeric"`
	RestaurantID int64  `json:"restaurant_id" binding:"required,min=1"`
	DishName     string `json:"dish_name" binding:"required"`
}

func (s *GinServer) validUser(ctx *gin.Context, userID int64, price float64, tx *gorm.DB) (*model.User, bool) {
	userData, err := s.store.GetUserStore().WithUserTx(tx).GetByID(ctx, userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return nil, false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return nil, false
	}
	if userData.CashBalance < price {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("cashBalance is not enough , userId %d", userID)})
		return nil, false
	}
	err = s.store.GetUserStore().WithUserTx(tx).DecreaseCashBalance(ctx, userData, price)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return nil, false
	}
	return userData, true
}

func (s *GinServer) validRestaurant(ctx *gin.Context, restaurantID int64, dishName string, tx *gorm.DB) (*model.Restaurant, float64, bool) {
	restaurantData, err := s.store.GetRestaurantStore().WithRestaurantTx(tx).GetByID(ctx, restaurantID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return nil, 0, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return nil, 0, false
	}

	for _, menu := range restaurantData.Menu {
		if menu.DishName == dishName {
			err = s.store.GetRestaurantStore().WithRestaurantTx(tx).IncreaseCashBalance(ctx, restaurantData, menu.Price)
			if err != nil {
				return nil, 0, false
			}
			return restaurantData, menu.Price, true
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "dish not found in restaurant"})
	return nil, 0, false
}

func (s *GinServer) createPurchaseHistory(ctx *gin.Context, history model.PurchaseHistory, tx *gorm.DB) error {
	_, err := s.store.GetHistoryStore().WithHistoryStoreTx(tx).Add(ctx, history)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return err
	}
	return nil
}

// createOrder
// @Summary      createOrder
// @Description  Process a user purchasing a dish from a restaurant
// @Accept       json
// @Produce      json
// @Param        data  body      purchaseRequest  true  "Purchase Request with UserId min=0 , RestaurantID min = 1"
// @Success      200   {string}  string "{"msg": "Successfully Purchase Order"}"
// @Failure      400   {string}  string "{"err": "err string"}"
// @Router       /api/v1/purchase [POST]
func (s *GinServer) createOrder(ctx *gin.Context) {
	var req purchaseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	tx := ctx.MustGet("db_trx").(*gorm.DB)
	restaurantData, price, valid := s.validRestaurant(ctx, req.RestaurantID, req.DishName, tx)
	if !valid {
		tx.Rollback()
		return
	}
	userData, valid := s.validUser(ctx, *req.UserID, price, tx)
	if !valid {
		tx.Rollback()
		return
	}
	history := model.PurchaseHistory{
		UserId:            userData.ID,
		DishName:          req.DishName,
		RestaurantName:    restaurantData.RestaurantName,
		TransactionAmount: price,
		TransactionDate:   time.Now().Format("2006-01-02 15:04:05"),
	}
	err := s.createPurchaseHistory(ctx, history, tx)
	if err != nil {
		tx.Rollback()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully Purchase Order"})
}
