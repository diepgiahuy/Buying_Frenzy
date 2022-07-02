package main

import (
	"context"
	"github.com/diepgiahuy/Buying_Frenzy/infra"
	"github.com/urfave/cli"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type ApplicationContext struct {
	ctx context.Context
	db  *gorm.DB
}

func main() {
	ctx := context.Background()
	application, cleanup, err := InitApplication(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	handleSigterm(cleanup)
	err = application.Commands().Run(os.Args)
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

// Import creates a command that start an http server
func (a *ApplicationContext) Import() cli.Command {
	return cli.Command{
		Name:  "import",
		Usage: "import json data",
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}

func (a *ApplicationContext) Commands() *cli.App {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		a.Serve(),
		a.Import(),
	}
	return app
}

func run() {
	ginServer := infra.NewServer(
		8080,
		infra.DebugMode,
	)
	ginServer.Start()
}
