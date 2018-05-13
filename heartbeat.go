package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	// Read data file that specifies which sites to check:
	dataFilePath := os.Args[1]
	if dataFilePath == "" {
		log.Fatal("Configuration file path not specified in command line arguments")
	}
	data, err := loadData(dataFilePath)
	check(err)
	stethoscope(&data)
	fmt.Println(data)

	jsonData, err := json.Marshal(data)
	check(err)

	// Save data
	err = ioutil.WriteFile(dataFilePath, jsonData, 0644)
	check(err)
}

// stethoscope checks the heartbeats of each site in the records object
func stethoscope(data *records) {
	for i := range data.Records {
		// Get mem address of each record
		rec := &data.Records[i]
		// Check the heartbeat
		code, err := checkHeartbeat(rec.URL)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Save data
		rec.Heartbeats = append(rec.Heartbeats, heartbeat{time.Now().String(), code})
	}

}

// checkHeartbeat calls http.Get on url and returns the StatusCode and a possible error
func checkHeartbeat(url string) (int, error) {
	resp, err := http.Get(url)
	fmt.Printf("Checking %s: ", url)
	if err != nil {
		return 0, err
	}
	fmt.Printf("%s has code %s\n", url, resp.Status)
	return resp.StatusCode, nil

}

// loadData loads the save json data used to keep heartbeat records
func loadData(path string) (rec records, err error) {
	dat, err := ioutil.ReadFile(path)
	check(err)
	rec = records{}
	json.Unmarshal(dat, &rec)
	fmt.Println("Data loaded")
	return
}

type heartbeat struct {
	Date       string `json:"date"`
	StatusCode int    `json:"StatusCode"`
}
type record struct {
	URL        string      `json:"url"`
	Heartbeats []heartbeat `json:"heartbeats"`
}

type records struct {
	Records []record `json:"records"`
}
