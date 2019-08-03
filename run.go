package main

import (
    "os/exec"

    "github.com/urfave/cli"
)

func run(c *cli.Context) error {
    var cmd *exec.Cmd
    if NetworkType == "mainnet" {
        cmd = exec.Command("java",
            "-jar", nodeJar,
            "-c", configFile)
    } else {
        cmd = exec.Command("java",
            "-jar", nodeJar,
            "-c", configFile,
            "--witness")
    }
    err := cmd.Run()
    if err != nil {
        return err
    }
    return nil
}
