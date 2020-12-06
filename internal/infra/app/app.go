package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/hoangson1024/golang/internal/infra"
	db_test_service "github.com/hoangson1024/golang/pkg/test-db/service"

	"github.com/urfave/cli"
)

type ApplicationContext struct {
	ctx     context.Context
	cfg     *infra.AppConfig
	wrapper db_test_service.Wrapper
}

func (a *ApplicationContext) Commands() *cli.App {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		a.Serve(),
	}

	return app
}

// HandleSigterm -- Handles Ctrl+C or most other means of "controlled" shutdown gracefully.
// Invokes the supplied func before exiting.
func HandleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}
