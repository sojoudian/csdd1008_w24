package main

import "fmt"

const worldTimeAPI = "http://worldtimeapi.org/api/timezone/America/Toronto"

type TimeInfo struct {
	Datetime string `json:"datetime"`
}

func main() {
	fmt.Println("Server started!")
}
