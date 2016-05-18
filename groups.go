package main

import (
  "fmt"
  "os"
  "encoding/json"
  "github.com/nlopes/slack"
  "./config"
)

type Config struct {
  Key      string
  Channel  string
  Launcher string
}

func main() {

  config := Config.new("config.json")

  api := slack.New(config.Key)
  groups, err := api.GetGroups(false)
  if err != nil {
    fmt.Printf("%s\n", err)
    return
  }
  for _, group := range groups {
    fmt.Printf("%s %s\n", group.Name, group.ID)
  }
}
