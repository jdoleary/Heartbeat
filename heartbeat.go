package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	// Read data file that specifies which sites to check:
	if len(os.Args) != 2 {
		log.Fatal("ERR: dataDirPath not specified in command line arguments")
	}
	dataDirPath := os.Args[1]
	if dataDirPath == "" {
		log.Fatal("ERR: dataDirPath not specified in command line arguments")
	}
	var builder strings.Builder
	builder.WriteString(dataDirPath)
	builder.WriteString("/data.json")
	dataFilePath := builder.String()
	data, err := loadData(dataFilePath)
	check(err)
	stethoscope(&data)

	jsonData, err := json.Marshal(data)
	check(err)

	// Save data
	err = ioutil.WriteFile(dataFilePath, jsonData, 0644)
	check(err)

	// Save pretty data
	var builderPretty strings.Builder
	builderPretty.WriteString(dataDirPath)
	builderPretty.WriteString("/pretty.txt")
	dataPrettyFilePath := builderPretty.String()
	prettyData := prettifyRecords(&data)
	err = ioutil.WriteFile(dataPrettyFilePath, []byte(prettyData), 0644)
	check(err)
}

func prettifyRecords(data *records) string {
	var builder strings.Builder
	for i := range data.Records {
		// Get mem address of each record
		rec := &data.Records[i]
		builder.WriteString("\n")
		builder.WriteString("\n")
		builder.WriteString(rec.URL)
		builder.WriteString("\n")
		for _, hb := range rec.Heartbeats {
			builder.WriteString("\t")
			builder.WriteString(hb.Date)
			builder.WriteString(" | ")
			builder.WriteString(strconv.Itoa(hb.StatusCode))
			builder.WriteString("\n")
		}

	}
	return builder.String()
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
		rec.Heartbeats = append(rec.Heartbeats, heartbeat{time.Now().Format("Mon Jan _2 2006 15:04:05"), code})
		// Limit heartbeat records
		if len(rec.Heartbeats) > heartbeatsMaxLength {
			rec.Heartbeats = rec.Heartbeats[len(rec.Heartbeats)-heartbeatsMaxLength:]
		}
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

var heartbeatsMaxLength = 24
