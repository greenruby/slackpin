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
  pins, _, err := api.ListPins(config.Channel)
  if err != nil {
    fmt.Printf("%s\n", err)
    return
  }
  for _, pin := range pins {
    fmt.Println(pin.Message.Text)
  }
}
