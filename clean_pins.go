package main

import (
  "os"
  "fmt"
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
    fmt.Printf("Error listing pins: %s\n", err)
    return
  }
  for _, pin := range pins {
    msgRef := slack.NewRefToMessage(config.Channel, pin.Message.Timestamp)
    // fmt.Printf("%+v\n", msgRef)
    err = api.RemovePin(config.Channel, msgRef)
    if err != nil {
      fmt.Printf("Error remove pin: %s\n", err)
      return
    }
    fmt.Printf("\"%s\" removed.\n", pin.Message.Text)
  }
}
