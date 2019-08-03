package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	NetworkType  string
	FullNodePort string
	SolNodePort  string
)

func main() {
	app := cli.NewApp()
	app.Name = "troncli"
	app.Version = "0.0.2"
	app.Usage = "go version tron-cli"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "typen, tn",
			Usage:       "network type",
			Value:       "mainnet",
			Destination: &NetworkType,
		},
		cli.StringFlag{
			Name:        "portf, pf",
			Usage:       "fullnode port",
			Value:       "8090",
			Destination: &FullNodePort,
		},
		cli.StringFlag{
			Name:        "ports, ps",
			Usage:       "solidity port",
			Value:       "8091",
			Destination: &SolNodePort,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init",
			Action:  initialize,
		},
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "config",
			Action:  configure,
		},
		{
			Name:    "run",
			Aliases: []string{"c"},
			Usage:   "run",
			Action:  run,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
