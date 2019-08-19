package main

import (
    "fmt"
    "io"
    "net/http"
    "os"

    "github.com/urfave/cli"
)

func initialize(c *cli.Context) error {
    fmt.Fprintf(c.App.Writer, "Downloading files...\n")
    if err := download(c, releaseUrl, nodeJar); err != nil {
        return err
    }
    if err := download(c, templateUrl, configTemplate); err != nil {
        return err
    }
    return nil
}

func download(c *cli.Context, url string, dest string) error {
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
    fmt.Fprintf(c.App.Writer,
        fmt.Sprintf("Successfully download %v from %v of size: %v KB \n",
            dest, url, size/1024),
    )

    return nil
}
