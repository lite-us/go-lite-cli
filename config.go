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
	nc.SeedIpList = `[
      "54.236.37.243:18888",
      "52.53.189.99:18888",
      "18.196.99.16:18888",
      "34.253.187.192:18888",
      "52.56.56.149:18888",
      "35.180.51.163:18888",
      "54.252.224.209:18888",
      "18.228.15.36:18888",
      "52.15.93.92:18888",
      "34.220.77.106:18888",
      "13.127.47.162:18888",
      "13.124.62.58:18888",
      "35.182.229.162:18888",
      "18.209.42.127:18888",
      "3.218.137.187:18888",
      "34.237.210.82:18888"
    ]`
	nc.LocalWitness = `[]`
	nc.Committee = `{allowCreationOfContracts = 0}`
	nc.BlockSyncCheck = "true"
	nc.MinParticipationRate = "15"
	nc.ConnectFactor = "0.3"
	nc.ActiveConnectFactor = "0.1"
	nc.GenesisBlock = `{
    // Reserve balance
    assets = [
      {
        accountName = "Zion"
        accountType = "AssetIssue"
        address = "TLLM21wteSPs4hKjbxgmH1L6poyMjeTbHm"
        balance = "99000000000000000"
      },
      {
        accountName = "Sun"
        accountType = "AssetIssue"
        address = "TXmVpin5vq5gdZsciyyjdZgKRUju4st1wM"
        balance = "0"
      },
      {
        accountName = "Blackhole"
        accountType = "AssetIssue"
        address = "TLsV52sRDL79HXGGm9yzwKibb6BeruhUzy"
        balance = "-9223372036854775808"
      }
    ]

    witnesses = [
      {
        address: THKJYuUmMKKARNf7s2VT51g5uPY6KEqnat,
        url = "http://GR1.com",
        voteCount = 100000026
      },
      {
        address: TVDmPWGYxgi5DNeW8hXrzrhY8Y6zgxPNg4,
        url = "http://GR2.com",
        voteCount = 100000025
      },
      {
        address: TWKZN1JJPFydd5rMgMCV5aZTSiwmoksSZv,
        url = "http://GR3.com",
        voteCount = 100000024
      },
      {
        address: TDarXEG2rAD57oa7JTK785Yb2Et32UzY32,
        url = "http://GR4.com",
        voteCount = 100000023
      },
      {
        address: TAmFfS4Tmm8yKeoqZN8x51ASwdQBdnVizt,
        url = "http://GR5.com",
        voteCount = 100000022
      },
      {
        address: TK6V5Pw2UWQWpySnZyCDZaAvu1y48oRgXN,
        url = "http://GR6.com",
        voteCount = 100000021
      },
      {
        address: TGqFJPFiEqdZx52ZR4QcKHz4Zr3QXA24VL,
        url = "http://GR7.com",
        voteCount = 100000020
      },
      {
        address: TC1ZCj9Ne3j5v3TLx5ZCDLD55MU9g3XqQW,
        url = "http://GR8.com",
        voteCount = 100000019
      },
      {
        address: TWm3id3mrQ42guf7c4oVpYExyTYnEGy3JL,
        url = "http://GR9.com",
        voteCount = 100000018
      },
      {
        address: TCvwc3FV3ssq2rD82rMmjhT4PVXYTsFcKV,
        url = "http://GR10.com",
        voteCount = 100000017
      },
      {
        address: TFuC2Qge4GxA2U9abKxk1pw3YZvGM5XRir,
        url = "http://GR11.com",
        voteCount = 100000016
      },
      {
        address: TNGoca1VHC6Y5Jd2B1VFpFEhizVk92Rz85,
        url = "http://GR12.com",
        voteCount = 100000015
      },
      {
        address: TLCjmH6SqGK8twZ9XrBDWpBbfyvEXihhNS,
        url = "http://GR13.com",
        voteCount = 100000014
      },
      {
        address: TEEzguTtCihbRPfjf1CvW8Euxz1kKuvtR9,
        url = "http://GR14.com",
        voteCount = 100000013
      },
      {
        address: TZHvwiw9cehbMxrtTbmAexm9oPo4eFFvLS,
        url = "http://GR15.com",
        voteCount = 100000012
      },
      {
        address: TGK6iAKgBmHeQyp5hn3imB71EDnFPkXiPR,
        url = "http://GR16.com",
        voteCount = 100000011
      },
      {
        address: TLaqfGrxZ3dykAFps7M2B4gETTX1yixPgN,
        url = "http://GR17.com",
        voteCount = 100000010
      },
      {
        address: TX3ZceVew6yLC5hWTXnjrUFtiFfUDGKGty,
        url = "http://GR18.com",
        voteCount = 100000009
      },
      {
        address: TYednHaV9zXpnPchSywVpnseQxY9Pxw4do,
        url = "http://GR19.com",
        voteCount = 100000008
      },
      {
        address: TCf5cqLffPccEY7hcsabiFnMfdipfyryvr,
        url = "http://GR20.com",
        voteCount = 100000007
      },
      {
        address: TAa14iLEKPAetX49mzaxZmH6saRxcX7dT5,
        url = "http://GR21.com",
        voteCount = 100000006
      },
      {
        address: TBYsHxDmFaRmfCF3jZNmgeJE8sDnTNKHbz,
        url = "http://GR22.com",
        voteCount = 100000005
      },
      {
        address: TEVAq8dmSQyTYK7uP1ZnZpa6MBVR83GsV6,
        url = "http://GR23.com",
        voteCount = 100000004
      },
      {
        address: TRKJzrZxN34YyB8aBqqPDt7g4fv6sieemz,
        url = "http://GR24.com",
        voteCount = 100000003
      },
      {
        address: TRMP6SKeFUt5NtMLzJv8kdpYuHRnEGjGfe,
        url = "http://GR25.com",
        voteCount = 100000002
      },
      {
        address: TDbNE1VajxjpgM5p7FyGNDASt3UVoFbiD3,
        url = "http://GR26.com",
        voteCount = 100000001
      },
      {
        address: TLTDZBcPoJ8tZ6TTEeEqEvwYFk2wgotSfD,
        url = "http://GR27.com",
        voteCount = 100000000
      }
    ]

    timestamp = "0" //2017-8-26 12:00:00

    parentHash = "0xe58f33f9baf9305dc6f82b9f1934ea8f0ade2defb951258d50167028c780351f"
  }`
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
	nc.GenesisBlock = `{
  # Reserve balance
  assets = [
    {
      accountName = "Zion"
      accountType = "AssetIssue"
      address = "TPL66VK2gCXNCD7EJg9pgJRfqcRazjhUZY"
      balance = "95000000000000000"
    },
    {
      accountName = "Sun"
      accountType = "AssetIssue"
      address = "TWsm8HtU2A5eEzoT8ev8yaoFjHsXLLrckb"
      balance = "5000000000000000"
    },
    {
      accountName = "Blackhole"
      accountType = "AssetIssue"
      address = "TSJD5rdu6wZXP7F2m3a3tn8Co3JcMjtBip"
      balance = "-9223372036854775808"
    },
    {
      accountName = "TestA"
      accountType = "AssetIssue"
      address = "TVdyt1s88BdiCjKt6K2YuoSmpWScZYK1QF"
      balance = "1000000000000000"
    },
    {
      accountName = "TestB"
      accountType = "AssetIssue"
      address = "TCNVmGtkfknHpKSZXepZDXRowHF7kosxcv"
      balance = "1000000000000000"
    },
    {
      accountName = "TestC"
      accountType = "AssetIssue"
      address = "TAbzgkG8p3yF5aywKVgq9AaAu6hvF2JrVC"
      balance = "1000000000000000"
    },
    {
      accountName = "TestD"
      accountType = "AssetIssue"
      address = "TMmmvwvkBPBv3Gkw9cGKbZ8PLznYkTu3ep"
      balance = "1000000000000000"
    },
    {
      accountName = "TestE"
      accountType = "AssetIssue"
      address = "TBJHZu4Sm86aWHtt6VF6KQSzot8vKTuTKx"
      balance = "1000000000000000"
    }
  ]

  witnesses = [
    {
      address: TPL66VK2gCXNCD7EJg9pgJRfqcRazjhUZY,
      url = "http://tronstudio.com",
      voteCount = 10000
    }
  ]

  timestamp = "0" #2017-8-26 12:00:00

  parentHash = "957dc2d350daecc7bb6a38f3938ebde0a0c1cedafe15f0edae4256a2907449f6"
}`
}
