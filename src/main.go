package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	startMonitoring()

	sites := returnSites()

	for {
		for _, site := range sites {
			siteStatus := getSiteStatus(site)
			fmt.Println(site, " Status ->", siteStatus)
		}
		fmt.Println()
		time.Sleep(60 * time.Second)
	}
}

func startMonitoring() {
	fmt.Println("Hello")
	fmt.Println("Monitoring...")
}

func getSiteStatus(site string) int {
	response, _ := http.Get(site)
	return response.StatusCode
}

func returnSites() []string {
	return []string{"https://www.facebook.com/", "https://www.instagram.com/", "https://random-status-code.herokuapp.com"}
}
