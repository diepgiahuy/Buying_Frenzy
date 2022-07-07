package main

import (
	"context"
	"encoding/json"
	"github.com/diepgiahuy/Buying_Frenzy/infra"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/model"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/storage"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type ApplicationContext struct {
	ctx context.Context
	db  *storage.Repo
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ctx := context.Background()
	cli, cleanup, err := InitApplication(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	handleSigterm(cleanup)
	err = cli.Commands().Run(os.Args)
	if err != nil {
		log.Fatalln(err.Error())
	}

}
func handleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}

// Serve creates a command that start an http server
func (a *ApplicationContext) Serve() cli.Command {
	return cli.Command{
		Name:  "serve",
		Usage: "serve http request",
		Action: func(c *cli.Context) error {
			run()
			return nil
		},
	}
}

// LoadData create a command that load json file to db
func (a *ApplicationContext) LoadData() cli.Command {
	return cli.Command{
		Name:  "load",
		Usage: "import json data",
		Action: func(c *cli.Context) error {
			loadData(a)
			return nil
		},
	}
}

func (a *ApplicationContext) Commands() *cli.App {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		a.Serve(),
		a.LoadData(),
	}
	return app
}

func run() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ginServer := infra.NewServer(
		8080,
		infra.DebugMode,
	)
	ginServer.Start()
}

type RawData struct {
	CashBalance float64 `json:"cashBalance"`
	Menu        []struct {
		DishName string  `json:"dishName"`
		Price    float64 `json:"price"`
	} `json:"menu"`
	OpeningHours   string `json:"openingHours"`
	RestaurantName string `json:"restaurantName"`
}

func loadRestaurantData(app *ApplicationContext, jsonData []byte) {
	var data []model.Restaurant
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	err = app.db.AddRestaurant(app.ctx, data)
	if err != nil {
		return
	}
}

func loadMenuData(app *ApplicationContext, jsonData []byte) {
	var data []model.Menu
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	err = app.db.AddMenu(app.ctx, data)
	if err != nil {
		return
	}
}

func loadData(app *ApplicationContext) {
	// Let's first read the `restaurant_with_menu.json` file
	jsonData, err := ioutil.ReadFile("./restaurant_with_menu.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	loadRestaurantData(app, jsonData)
}
