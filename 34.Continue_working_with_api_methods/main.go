package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type HandleStruct struct{}

func (h HandleStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().UTC().Format(time.RFC3339)
	buffer := []byte(currentTime + "\n")
	fmt.Println(currentTime)
	n, err := w.Write(buffer)
	if n == 0 || err != nil {
		panic("Eror write!")
	}
	
}

type CurrentTime struct {
	DayOfWeek time.Weekday `json:"day_of_week"`
	DayToday  int          `json:"day_of_month"`
	Month     time.Month   `json:"month"`
	Year      int          `json:"year"`
	Hour      int          `json:"hour"`
	Minute    int          `json:"minute"`
	Second    int          `json:"second"`
}

func GetJsonCurrentTime(w http.ResponseWriter, r *http.Request) {
	timeNow := time.Now()

	currentTime := CurrentTime{
		DayOfWeek: timeNow.Weekday(),
		DayToday:  timeNow.Day(),
		Month:     timeNow.Month(),
		Year:      timeNow.Year(),
		Hour:      timeNow.Hour(),
		Minute:    timeNow.Minute(),
		Second:    timeNow.Second(),
	}
	newEndcoder := json.NewEncoder(w)
	w.Header().Set("Content-Type","application/json")

	err := newEndcoder.Encode(currentTime)
	if err != nil {
		panic(err)
	}
	fmt.Println(newEndcoder)
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("/GET", HandleStruct{})

	server := http.Server{
		Addr:         ":9000",
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 35 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	fmt.Println("Server is listening....!")

	mux.HandleFunc("/GET/JSON/", GetJsonCurrentTime)

	err := server.ListenAndServe()
	defer server.Close()
	if err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
}