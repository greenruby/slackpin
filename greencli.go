package main

import (
  "fmt"
  "os"

  "github.com/codegangsta/cli"
)

func main() {
  cli.NewApp().Run(os.Args)
}
