package main

import (
    "fmt"
    "io"
    "net/http"
    "os"

    "github.com/urfave/cli"
)

func initialize(c *cli.Context) error {
    releaseUrl := "https://github.com/tronprotocol/java-tron/releases/download/Odyssey-v3.6.1/FullNode.jar"
    if err := download(releaseUrl, "FullNode.jar"); err != nil {
        return err
    }
    templateUrl := "https://raw.githubusercontent.com/TRON-US/go-tron-cli/master/config.tmpl"
    if err := download(templateUrl, "config.tmpl"); err != nil {
        return err
    }
    return nil
}

func download(url string, dest string) error {
    out, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer out.Close()

    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    size, err := io.Copy(out, resp.Body)
    fmt.Println("Download size: ", size)
    return nil
}
