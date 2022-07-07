package main

import (
	"context"
	"github.com/diepgiahuy/Buying_Frenzy/cmd"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ctx := context.Background()
	cli, cleanup, err := cmd.InitApplication(ctx)
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
