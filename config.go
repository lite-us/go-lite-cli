package main

import (
	"os"
	"text/template"

	"github.com/urfave/cli"
)

type NodeConfig struct {
	FullNodePort         string
	SolNodePort          string
	P2PVersion           string
	SeedIpList           string
	GenesisBlock         string
	LocalWitness         string
	BlockSyncCheck       string
	MinParticipationRate string
	ConnectFactor        string
	ActiveConnectFactor  string
	Committee            string
}

var nc NodeConfig

func configure(c *cli.Context) error {
	nc.FullNodePort = FullNodePort
	nc.SolNodePort = SolNodePort

	if NetworkType == "mainnet" {
		setMainNet()
	} else {
		setPrivateNet()
	}

	if err := writeConfigFile(); err != nil {
		return err
	}
	return nil
}

func writeConfigFile() error {
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

func setMainNet() {
	nc.P2PVersion = "11111"
	nc.SeedIpList = mainNetSeedIpList
	nc.LocalWitness = `[]`
	nc.Committee = `{allowCreationOfContracts = 0}`
	nc.BlockSyncCheck = "true"
	nc.MinParticipationRate = "15"
	nc.ConnectFactor = "0.3"
	nc.ActiveConnectFactor = "0.1"
	nc.GenesisBlock = mainNetGenesisBlock
}

func setPrivateNet() {
	nc.P2PVersion = "1"
	nc.SeedIpList = `[]`
	nc.LocalWitness = `[da146374a75310b9666e834ee4ad0866d6f4035967bfc76217c5a495fff9f0d0]`
	nc.BlockSyncCheck = "false"
	nc.Committee = `{allowCreationOfContracts = 1}`
	nc.MinParticipationRate = "0"
	nc.ConnectFactor = "0"
	nc.ActiveConnectFactor = "0"
	nc.GenesisBlock = privateNetGenesisBlock
}
