package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Goalt/go-template/cmd/subcomands"
	"github.com/Goalt/go-template/cmd/variables"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "go-template",
		Usage:    "./go-template",
		Commands: subcomands.Get(),
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    variables.DebugLevel,
				Value:   1,
				EnvVars: []string{variables.DebugLevel},
			},
			&cli.IntFlag{
				Name:    variables.MaxFileSize,
				Value:   1,
				EnvVars: []string{variables.MaxFileSize},
			},
			&cli.StringFlag{
				Name:    variables.RootPath,
				EnvVars: []string{variables.RootPath},
			},
			&cli.StringFlag{
				Name:    variables.SecretKey,
				EnvVars: []string{variables.SecretKey},
			},
			&cli.StringFlag{
				Name:    variables.MysqlDatabaseName,
				EnvVars: []string{variables.MysqlDatabaseName},
			},
			&cli.StringFlag{
				Name:    variables.MysqlUser,
				EnvVars: []string{variables.MysqlUser},
			},
			&cli.StringFlag{
				Name:    variables.MysqlPassword,
				EnvVars: []string{variables.MysqlPassword},
			},
			&cli.StringFlag{
				Name:    variables.MysqlHost,
				EnvVars: []string{variables.MysqlHost},
			},
			&cli.StringFlag{
				Name:    variables.MysqlPort,
				EnvVars: []string{variables.MysqlPort},
			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Println(os.Environ())

			for now := range time.Tick(time.Second * 4) {
				fmt.Println(now)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

type simpleInterface interface {
	SimpleInterfaceFunc(int) string
}

func simpleFunc(i simpleInterface) {
	log.Println(i.SimpleInterfaceFunc(42))
}
