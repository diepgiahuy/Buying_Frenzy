package util

import (
	"encoding/json"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"log"
)

func LoadUserData(jsonData []byte) []model.User {
	var userData []model.User
	err := json.Unmarshal(jsonData, &userData)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return userData
}
