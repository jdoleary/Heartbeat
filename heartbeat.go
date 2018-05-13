package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("|||Heartbeat|||\n")
	checkHeartbeat("https://jdoleary.me")
}

func checkHeartbeat(url string) {
	resp, err := http.Get(url)
	fmt.Printf("Checking %s\n", url)
	if err != nil {
		log.Fatal(err);
	}
	fmt.Printf("%s is alive with code %s\n", url, resp.Status)

}
