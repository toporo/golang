package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
	response, err := http.Get(site)
	if err != nil {
		fmt.Printf("Error found on gite site", site, ": ", err)
		return 400
	}

	return response.StatusCode
}

func returnSites() []string {
	file, err := os.Open("../resource/sites.txt")
	readerUtil := bufio.NewReader(file)

	if err != nil {
		fmt.Println("Error found: ", err)
	}

	var sites []string

	for {
		line, err := readerUtil.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}
