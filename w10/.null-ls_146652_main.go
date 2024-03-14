package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const worldTimeAPI = "http://worldtimeapi.org/api/timezone/America/Toronto"

type TimeInfo struct {
	Datetime string `json:"datetime"`
}

func getTorontoTime() (string, error) {
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

	return TimeInfo.Datetime, nil
}

func torontoTimeHandler(w http.ResponseWriter, r *http.Request) {
	torontoTime, err := getTorontoTime()
	if err != nil {
		http.Error(w, "Error fetchin Toronto time", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Toronto time is %s", torontoTime)
	resp := map[string]string{"current_Time_Toronto": torontoTime}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/torontotime", torontoTimeHandler)
	fmt.Println("Server is started on port 8014!")
	log.Fatal(http.ListenAndServe(":8014", nil))
}
