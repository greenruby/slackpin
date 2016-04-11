package main

import (
  "fmt"
  "os"
  "encoding/json"
  "github.com/nlopes/slack"
)

type Config struct {
  Key      string
  Channel  string
}


func main() {

  file, _ := os.Open("config.json")
  decoder := json.NewDecoder(file)
  config := Config{}
  err := decoder.Decode(&config)
  if err != nil {
    fmt.Println("error:", err)
    return
  }

  api := slack.New(config.Key)
  channels, err := api.GetChannels(false)
  if err != nil {
    fmt.Printf("%s\n", err)
    return
  }
  for _, channel := range channels {
    fmt.Printf("%s %s\n", channel.Name, channel.ID)
  }
}
