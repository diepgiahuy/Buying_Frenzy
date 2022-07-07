package utli

import (
	"encoding/json"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"log"
)

type rawUserData struct {
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

func LoadUserData(jsonData []byte) []model.User {
	var userData []model.User
	err := json.Unmarshal(jsonData, &userData)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return nil
}
