package user

type Model struct {
	CashBalance     float64 `json:"cashBalance"`
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	PurchaseHistory []struct {
		DishName          string  `json:"dishName"`
		RestaurantName    string  `json:"restaurantName"`
		TransactionAmount float64 `json:"transactionAmount"`
		TransactionDate   string  `json:"transactionDate"`
	} `json:"purchaseHistory"`
}
