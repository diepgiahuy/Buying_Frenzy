package cmd

import (
	"context"
	"github.com/diepgiahuy/Buying_Frenzy/infra"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/storage"
	"github.com/diepgiahuy/Buying_Frenzy/util"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
)

type ApplicationContext struct {
	Ctx context.Context
	Db  *storage.Repo
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
func (app *ApplicationContext) LoadData() cli.Command {
	return cli.Command{
		Name:  "load",
		Usage: "import json data",
		Action: func(c *cli.Context) error {
			//loadRestaurantData(app)
			loadUserData(app)
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

func loadRestaurantData(app *ApplicationContext) {
	// Let's first read the `restaurant_with_menu.json` file
	jsonData, err := ioutil.ReadFile("./restaurant_with_menu.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	restaurants := util.LoadRestaurantData(jsonData)
	err = app.Db.AddRestaurantWithBatches(app.Ctx, restaurants)
	if err != nil {
		return
	}
}

func loadUserData(app *ApplicationContext) {
	// Let's second read the `users_with_purchase_history.json` file
	jsonData, err := ioutil.ReadFile("./users_with_purchase_history.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	users := util.LoadUserData(jsonData)
	err = app.Db.AddUserWithBatches(app.Ctx, users)
	if err != nil {
		return
	}
}
