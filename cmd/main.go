package main

import (
	"context"
	"log"
	"os"

	"github.com/hoangson1024/golang/internal/infra/app"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	cli, cleanup, err := app.InitApplication(ctx)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	app.HandleSigterm(cleanup)

	err = cli.Commands().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
