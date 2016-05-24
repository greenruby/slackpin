package main

import (
  // "fmt"
  "os"
  "gopkg.in/urfave/cli.v2"
)

func main() {
  app := cli.NewApp()
  
  app.Name = "greencli"
  app.Version = "0.0.1"
  app.Usage = "A multipurpose CLI for Green Ruby publication system."
  
  app.Commands = []cli.Command{
    Name:         "pin",
    Usage:        "work with Slack pins",
    // Aliases:      []string{"p"},
    // Subcommands:  []cli.Command{
    //   {
    //     Name:     "list",
    //     Usage:    "List all the pins",
    //     Aliases:  []string{"l"},
    //     Action:   func(c *cli.Context) error {
    //       fmt.Println("new task template: ", c.Args().First())
    //       return nil
    //     },
    //   },
    // },
  }

  app.Run(os.Args)
}
