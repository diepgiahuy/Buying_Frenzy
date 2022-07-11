package api

import (
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type listRestaurantRequest struct {
	Date string `form:"date" binding:"datetime=2006-01-02 15:04:05"` // Date with format 2006-01-02 15:04:05
	paginationRequest
}

type listRestaurantWithDishesRequest struct {
	TopList    *int     `form:"top_list" binding:"required,min=0"`                              // For get Top Y restaurants
	HighPrice  *float32 `form:"high_price,omitempty" binding:"required,min=0,gtfield=LowPrice"` // Higher price min = 0 and greater than Lower Price
	LowPrice   *float32 `form:"low_price,omitempty" binding:"required,min=0"`                   // Lower price min = 0
	Comparison *int     `form:"comparison,omitempty" binding:"required,min=0,max=1"`            // Param to get more of less dish to compare, With 0 is more and 1 is less otherwise throw error
	NumDishes  *int     `form:"num_dishes" binding:"required,min=0"`                            // Param to find number of dishes in restaurant within a price range
}

type listRestaurantByNameRequest struct {
	Name string `uri:"name" binding:"required"` // Name need to find
}

type paginationRequest struct {
	PageID   int `form:"page_id" binding:"required,min=1"`          // Page wants to get
	PageSize int `form:"page_size" binding:"required,min=5,max=10"` // Number of data in page min is 5 and max is 10
}

// listRestaurantsOpen
// @Summary      listRestaurantsOpen
// @Description  Get list restaurant open at certain date time
// @Produce      json
// @Param        restaurant  query     listRestaurantRequest true "listRestaurantsOpen"
// @Success      200  	     {object}  []model.Restaurant
// @Failure      400 		 {string}  string "{"err": "err string"}"
// @Router       /api/v1/restaurants [get]
func (s *GinServer) listRestaurantsOpen(ctx *gin.Context) {
	var req listRestaurantRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	offset := (req.PageID - 1) * req.PageSize
	res, err := s.store.GetRestaurantStore().GetRestaurantByDate(ctx, req.Date, offset, req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}

// listRestaurantsWithComparison
// @Summary      listRestaurantsWithComparison
// @Description  List top y restaurants that have more or less than x number of dishes within a price range, ranked alphabetically.
// @Produce      json
// @Param        restaurant  query      listRestaurantWithDishesRequest  true  "listRestaurantsWithComparison"
// @Success      200  	     {object}  []model.Restaurant
// @Failure 400 {string}     string "{"err": "err string"}"
// @Router       /api/v1/restaurants/top-list-with-price [get]
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

// listRestaurantsByName
// @Summary      listRestaurantsByName
// @Description  Search for restaurants by name, ranked by relevance to search term.
// @Produce      json
// @Param        restaurant     query    paginationRequest  true  "listRestaurantsByName"
// @Param        restaurant     path     string  true  "Restaurant Name"
// @Success      200  	{object}  []model.Restaurant
// @Failure      400 {string} string "{"err": "err string"}"
// @Router       /api/v1/restaurants/{name} [get]
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

// listDishByName
// @Summary      listDishByName
// @Description  Search for dish by name, ranked by relevance to search term.
// @Produce      json
// @Param        restaurant     query    paginationRequest  true  "listDishByName"
// @Param        restaurant     path     string  true  "Dish Name"
// @Success      200  	{object}  []model.Menu
// @Failure 400 {string} string "{"err": "err string"}"
// @Router       /api/v1/restaurants/dish/{name} [get]
func (s *GinServer) listDishByName(ctx *gin.Context) {
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
