package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
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

			registerLogs(site, siteStatus)
		}
		fmt.Println()
		time.Sleep(15 * time.Second)
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

func registerLogs(site string, status int) {
	file, err := os.OpenFile("../logs/logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + " - Status: " + strconv.Itoa(status) + "\n")

	file.Close()
}
