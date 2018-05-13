package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./data/heartbeats.json")
	check(err)
	fmt.Printf("%s\n", dat)
	rec := records{}
	json.Unmarshal(dat, &rec)
	fmt.Println(rec)

	// // Read conf file that specifies which sites to check:
	// confFilePath := os.Args[1]
	// if confFilePath == "" {
	// 	log.Fatal("Configuration file path not specified in command line arguments")
	// }
	// fmt.Printf("|||Heartbeat|||\n%s conf file loaded\n", confFilePath)
	// dat, err := ioutil.ReadFile(confFilePath)
	// check(err)
	// fileData := string(dat[:])
	// // Remove \r
	// fileData = strings.Replace(fileData, "\r", "", -1)
	// // Split by newline
	// sites := strings.Split(fileData, "\n")

	// // Loop and check sites:
	// for _, el := range sites {
	// 	checkHeartbeat(el)
	// }
}

// checkHeartbeat calls http.Get on url and returns the StatusCode and a possible error
func checkHeartbeat(url string) (int, error) {
	resp, err := http.Get(url)
	fmt.Printf("Checking %s: ", url)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Printf("%s has code %s\n", url, resp.Status)
	return resp.StatusCode, nil

}

type heartbeat struct {
	Date       string `json:"date"`
	StatusCode int    `json:"StatusCode"`
}
type record struct {
	URL        string      `json:url`
	Heartbeats []heartbeat `json:"heartbeats"`
}

type records struct {
	Records []record `json:"records"`
}
