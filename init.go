package main

import (
    "fmt"
    "io"
    "net/http"
    "os"

    "github.com/urfave/cli"
)

func initialize(c *cli.Context) error {
    if err := download(releaseUrl, nodeJar); err != nil {
        return err
    }
    if err := download(templateUrl, configTemplate); err != nil {
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
