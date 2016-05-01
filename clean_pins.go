package main

import (
  "os"
  "fmt"
  "regexp"
  "encoding/json"
  "github.com/nlopes/slack"
)

type Config struct {
  Key      string
  Channel  string
  Launcher string
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
  pins, _, err := api.ListPins(config.Channel)
  if err != nil {
    fmt.Printf("%s\n", err)
    return
  }
  for _, pin := range pins {
    _, err = api.RemovePin(config.Channel, pin)
    if err != nil {
      fmt.Printf("%s\n", err)
      return
    }
    fmt.Printf("%s removed.\n", pin)
  }
}
