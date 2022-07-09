package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type listRestaurantRequest struct {
	PageID   int    `form:"page_id" binding:"required,min=1"`
	PageSize int    `form:"page_size" binding:"required,min=5,max=10"`
	Date     string `form:"date" binding:"datetime=2006-01-02 15:04:05"`
}

type listRestaurantWithDishesRequest struct {
	Top       *int     `form:"top" binding:"required,min=0"`
	PriceTop  *float32 `form:"price_top,omitempty" binding:"required,min=0" `
	PriceBot  *float32 `form:"price_bot,omitempty" binding:"required,min=0"`
	NumDishes *int     `form:"num_dishes" binding:"required,min=0"`
}

type listRestaurantUriRequest struct {
	Name string `uri:"name" binding:"required"`
}

type listRestaurantByNameRequest struct {
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
	res, err := s.store.GetRestaurantWithDate(ctx, req.Date, offset, req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *GinServer) listRestaurantsWithMoreDishes(ctx *gin.Context) {
	var req listRestaurantWithDishesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	res, err := s.store.GetRestaurantWithMoreDishes(ctx, *req.PriceBot, *req.PriceTop, *req.NumDishes, *req.Top)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *GinServer) listRestaurantsWithLessDishes(ctx *gin.Context) {
	var req listRestaurantWithDishesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	res, err := s.store.GetRestaurantWithLessDishes(ctx, *req.PriceBot, *req.PriceTop, *req.NumDishes, *req.Top)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *GinServer) listRestaurantsByName(ctx *gin.Context) {
	var req listRestaurantByNameRequest
	var uri listRestaurantUriRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	offset := (req.PageID - 1) * req.PageSize
	res, err := s.store.GetRestaurantByTerm(ctx, uri.Name, offset, req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (s *GinServer) listRestaurantsByDishName(ctx *gin.Context) {
	var req listRestaurantByNameRequest
	var uri listRestaurantUriRequest
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	offset := (req.PageID - 1) * req.PageSize
	res, err := s.store.GetRestaurantByDishTerm(ctx, uri.Name, offset, req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
