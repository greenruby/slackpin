package main

import (
  "os"
  "fmt"
  "regexp"
  "strings"
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
  re := regexp.MustCompile(".*<([^>]*)>.*")
  users := make(map[string]string)
  for _, pin := range pins {
    url := fmt.Sprintf(re.ReplaceAllString(pin.Message.Text, "${1}"))
    if (users[pin.Message.User] == "") {
      user, err := api.GetUserInfo(pin.Message.User)
      if err != nil {
        fmt.Printf("%s\n", err)
        return
      }
      users[pin.Message.User] = user.Name
    }
    fmt.Println(config.Launcher, url)
  }
  i := 0
  userlist := make([]string, len(users))
  for _, v := range users {
    userlist[i] = v
    i++
  }
  fmt.Println(strings.Join(userlist[:], ", "))
}
