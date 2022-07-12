package storage

import (
	"github.com/diepgiahuy/Buying_Frenzy/util/config"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
	"testing"
)

var userStore *UserStore
var restaurantStore *RestaurantStore

var menuStore *MenuStore
var historyStore *HistoryStore
var operationHourStore *OperationHourStore

func TestMain(t *testing.M) {
	err := godotenv.Load("../../util/config/dev.env")
	if err != nil {
		log.Fatalf("Error loading .env file when test")
	}
	var cfg = config.Config{}
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Error processing .env file when test")
	}
	db := NewDB(cfg.Host, cfg.User, cfg.Password, cfg.Db, cfg.PostgresConfig.Port)
	userStore = db.GetUserStore()
	restaurantStore = db.GetRestaurantStore()
	menuStore = db.GetMenuStore()
	historyStore = db.GetHistoryStore()
	operationHourStore = db.GetOperationHourStore()

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	os.Exit(t.Run())
}
