package main

import (
    "fmt"
    "os/exec"

    "github.com/urfave/cli"
)

func run(c *cli.Context) error {
    var cmd *exec.Cmd
    if NetworkType == "mainnet" {
        cmd = exec.Command("java",
            "-jar", nodeJar,
            "-c", configFile, "&")
    } else {
        cmd = exec.Command("java",
            "-jar", nodeJar,
            "-c", configFile,
            "--witness", "&")
    }
    fmt.Fprintf(c.App.Writer,
        fmt.Sprintf(
            "Fullnode API - http://localhost:%v \nSolidity API - http://localhost:%v \n ",
            FullNodePort,
            SolNodePort,
        ),
    )
    err := cmd.Run()
    if err != nil {
        return err
    }
    return nil
}
