package config

import (
  "os"
  "fmt"
  "encoding/json"
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

}
