package util

import (
	"encoding/json"
	"fmt"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"log"
	"regexp"
	"strings"
	"time"
)

type rawRestaurantData struct {
	CashBalance    float64 `json:"cashBalance"`
	Menu           []model.Menu
	OpeningHours   string `json:"openingHours"`
	RestaurantName string `json:"restaurantName"`
}

func LoadRestaurantData(jsonData []byte) []model.Restaurant {
	var rawData []rawRestaurantData
	var restaurantsData []model.Restaurant
	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	for _, data := range rawData {
		restaurantData := model.Restaurant{
			CashBalance:    data.CashBalance,
			RestaurantName: data.RestaurantName,
			Menu:           data.Menu,
			OperationHour:  transformHourData(data.OpeningHours),
		}
		restaurantsData = append(restaurantsData, restaurantData)
	}
	return restaurantsData
}

func transformHourData(opsHour string) []model.OperationHour {
	regexFindHour := regexp.MustCompile(`((1[0-2]|0?[1-9])(?::([0-5][0-9]))? ?\s([AaPp][Mm]))`)
	regexFormatTime := regexp.MustCompile(`\s(1[0-2]|0?[1-9])(\s[AaPp][Mm])`)
	regexReplaceSpecialChar := regexp.MustCompile(`[-]`)
	regexRemoveSpace := regexp.MustCompile(`[\s]`)
	//"Mon, Fri 2:30 pm - 8 pm"
	//"Mon - Weds
	opsHour = regexFormatTime.ReplaceAllString(opsHour, " $1:00 $2")
	splitDate := strings.Split(opsHour, "/")
	formatToLayout := "3:04 pm"
	parseToLayout := "15:04:05"
	var operationHours []model.OperationHour
	for _, date := range splitDate {
		hour := regexFindHour.FindAllString(date, -1)
		date = regexFindHour.ReplaceAllString(date, "")
		date = regexReplaceSpecialChar.ReplaceAllString(date, ",")
		date = regexRemoveSpace.ReplaceAllString(date, "")
		dayOfWeek := strings.Split(date, ",")
		t, err := time.Parse(formatToLayout, hour[0])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		openHour := t.Format(parseToLayout)
		t, err = time.Parse(formatToLayout, hour[1])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		closeHour := t.Format(parseToLayout)
		for _, day := range dayOfWeek {
			if day == "" {
				continue
			}
			operationHour := model.OperationHour{
				Day:       day,
				OpenHour:  openHour,
				CloseHour: closeHour,
			}
			operationHours = append(operationHours, operationHour)

		}
	}
	return operationHours
}
