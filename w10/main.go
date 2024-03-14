package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const worldTimeAPI = "http://worldtimeapi.org/api/timezone/America/Toronto"

type TimeInfo struct{
  Datetime string `json:"datetime"`
}


func getTorontoTime() (string, error)  {
  resp, err := http.Get(worldTimeAPI)
  if err != nil {
    return "Error retriving data", err
  }
  defer resp.Body.Close()
  body, err := io.ReadAll(resp.Body)
  if err != nil {
    return "Error reading data", err
  }
  var TimeInfo TimeInfo
  err = json.Unmarshal(body, &TimeInfo)
  if err != nil {
    return "Error Parsing data", err
  }

  return TimeInfo.Datetime,  nil 
}

func main() {
	fmt.Println("Server started!")
}
