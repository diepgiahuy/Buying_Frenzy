package api

import (
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type listRestaurantRequest struct {
	Date string `form:"date" binding:"datetime=2006-01-02 15:04:05"`
	paginationRequest
}

type listRestaurantWithDishesRequest struct {
	TopList    *int     `form:"top_list" binding:"required,min=0"`
	HighPrice  *float32 `form:"high_price,omitempty" binding:"required,min=0,gtfield=LowPrice" `
	LowPrice   *float32 `form:"low_price,omitempty" binding:"required,min=0"`
	Comparison *int     `form:"comparison,omitempty" binding:"required,min=0,max=1"`
	NumDishes  *int     `form:"num_dishes" binding:"required,min=0"`
}

type listRestaurantByNameRequest struct {
	Name string `uri:"name" binding:"required"`
}

type paginationRequest struct {
	PageID   int `form:"page_id" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *GinServer) listRestaurantsOpen(ctx *gin.Context) {
	var req listRestaurantRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	offset := (req.PageID - 1) * req.PageSize
	res, err := s.store.GetRestaurantStore().GetRestaurantWithDate(ctx, req.Date, offset, req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *GinServer) listRestaurantsWithComparison(ctx *gin.Context) {
	var req listRestaurantWithDishesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var err error
	var res []model.Restaurant
	if *req.Comparison == 0 {
		res, err = s.store.GetRestaurantStore().GetRestaurantWithCompareMore(ctx, *req.LowPrice, *req.HighPrice, *req.NumDishes, *req.TopList)
	} else {
		res, err = s.store.GetRestaurantStore().GetRestaurantWithCompareLess(ctx, *req.LowPrice, *req.HighPrice, *req.NumDishes, *req.TopList)
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *GinServer) listRestaurantsByName(ctx *gin.Context) {
	var req paginationRequest
	var uri listRestaurantByNameRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	offset := (req.PageID - 1) * req.PageSize
	res, err := s.store.GetRestaurantStore().GetRestaurantByTerm(ctx, uri.Name, offset, req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *GinServer) listRestaurantsByDishName(ctx *gin.Context) {
	var req paginationRequest
	var uri listRestaurantByNameRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	offset := (req.PageID - 1) * req.PageSize
	res, err := s.store.GetRestaurantStore().GetRestaurantByDishTerm(ctx, uri.Name, offset, req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
