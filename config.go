package main

import (
	"os"
	"text/template"

	"github.com/urfave/cli"
)

type NodeConfig struct {
	Type         string
	FullNodePort string
	SolNodePort  string
}

func configure(c *cli.Context) error {
	var nc NodeConfig
	nc.Type = networkType
	nc.FullNodePort = fullNodePort
	nc.SolNodePort = solNodePort

	// load config template from file
	t, err := template.ParseFiles(configTemplate)
	if err != nil {
		return err
	}

	// create file descriptor
	f, err := os.Create(configFile)
	if err != nil {
		f.Close()
		return err
	}

	// write node config
	err = t.Execute(f, nc)
	if err != nil {
		return err
	}

	return nil
}
