package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/api"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/storage"
	"github.com/diepgiahuy/Buying_Frenzy/util"
	"github.com/diepgiahuy/Buying_Frenzy/util/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
)

type ApplicationContext struct {
	Ctx         context.Context
	Db          *storage.PostgresStore
	cfg         *config.Config
	httpHandler *api.GinServer
}

// Serve creates a command that start an http server
func (app *ApplicationContext) Serve() cli.Command {
	return cli.Command{
		Name:  "serve",
		Usage: "serve http request",
		Action: func(c *cli.Context) error {
			runDBMigration(*app.cfg)
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
func runDBMigration(cfg config.Config) {
	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.PostgresConfig.User, cfg.PostgresConfig.Password, cfg.PostgresConfig.Host, cfg.PostgresConfig.Port, cfg.PostgresConfig.Db)
	if os.Getenv("HEROKU_ENV") == "PROD" {
		dbSource = os.Getenv("DATABASE_URL")
	}
	db, err := sql.Open("postgres", dbSource)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		"postgres", driver)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("db migrated successfully")
}
