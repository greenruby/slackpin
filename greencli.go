package main

import (
  "fmt"
  "os"

  "github.com/codegangsta/cli"
)

func main() {
  app := cli.NewApp()
  app.Name = "greencli"
  app.Version = "0.0.1"
  app.Usage = "A multipurpose CLI for Green Ruby publication system."
  app.Action = func(c *cli.Context) error {
    fmt.Println("boom! I say!")
    return nil
  }

  app.Run(os.Args)
}
