package cmd

import (
	"context"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/api"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/storage"
	"github.com/diepgiahuy/Buying_Frenzy/util"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
)

type ApplicationContext struct {
	Ctx         context.Context
	Db          *storage.PostgresStore
	httpHandler *api.GinServer
}

// Serve creates a command that start an http server
func (app *ApplicationContext) Serve() cli.Command {
	return cli.Command{
		Name:  "serve",
		Usage: "serve http request",
		Action: func(c *cli.Context) error {
			run(app)
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
			loadRestaurantData(app)
			//loadUserData(app)
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

func run(app *ApplicationContext) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app.httpHandler.Start()
}

func loadRestaurantData(app *ApplicationContext) {
	// Let's first read the `restaurant_with_menu.json` file
	jsonData, err := ioutil.ReadFile("./restaurant_with_menu.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	restaurants := util.LoadRestaurantData(jsonData)
	err = app.Db.GetRestaurantStore().AddRestaurantWithBatches(app.Ctx, restaurants)
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
	err = app.Db.GetUserStore().AddUserWithBatches(app.Ctx, users)
	if err != nil {
		return
	}
}
