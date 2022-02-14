package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const BASE_DIR = "src/app"
const times_of_monitoring = 5
const delay_between_monitoring = 5 // Seconds

func main() {
	introduction()
	readJSONSites()
	for {
		showMenu()
		command := getCommandSelected()

		switch command {
		case 0:
			fmt.Println("Exit!")
			os.Exit(0)
		case 1:
			initMonitore()
		case 2:
			fmt.Println("Showing logs...")
		default:
			fmt.Println("Command not found! Please select a valid command.")
			os.Exit(-1)
		}
	}
}

func introduction() {
	name := "Laurinha"
	version := 1.1
	fmt.Println("Hello srta.", name)
	fmt.Println("This app is in version:", version)
}

func showMenu() {
	fmt.Println("1 - Init monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func getCommandSelected() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The command has selected was:", command)
	return command
}

type Site struct {
	Url         string `json:"url"`
	Environment string `json:"environment"`
	Project     string `json:"project"`
}

func initMonitore() {
	fmt.Println("Starting monitoring...")

	sites_to_monitore := readJSONSites()
	for i := 0; i < times_of_monitoring; i++ {
		fmt.Println("----------------------")
		fmt.Println("Testing", i+1, "time(s)")
		fmt.Println("----------------------")
		for _, site := range sites_to_monitore {
			testSiteStatus(site)
		}
		time.Sleep(delay_between_monitoring * time.Second)
	}
}

func testSiteStatus(site Site) {
	response, err := http.Get(site.Url)
	if err != nil {
		fmt.Println("An error occurred")
		fmt.Println(err)
	}
	if response.StatusCode == 200 {
		fmt.Println("SITE:", site.Url)
		fmt.Println("PROJECT:", site.Project)
		fmt.Println("STATUS", response.StatusCode)
		fmt.Println("ENVIRONMENT:", site.Environment)
		fmt.Println("Site is OK!")
	} else {
		fmt.Println("SITE:", site.Url)
		fmt.Println("PROJECT:", site.Project)
		fmt.Println("STATUS:", response.StatusCode)
		fmt.Println("ENVIRONMENT:", site.Environment)
		fmt.Println("Site", site, "have a problem :(")
	}
}

func readJSONSites() []Site {
	JSONFile := fmt.Sprintf("%s/sites-to-monitore.json", BASE_DIR)
	// file, err := os.Open(JSONFile)
	file, err := ioutil.ReadFile(JSONFile)
	if err != nil {
		fmt.Println("----------------------")
		fmt.Println("An error in OPEN occurred")
		fmt.Println(err)
		fmt.Println("----------------------")
	}
	var parsedJSON []Site
	err = json.Unmarshal(file, &parsedJSON)
	if err != nil {
		fmt.Println("----------------------")
		fmt.Println("An error in CONVERT JSON occurred")
		fmt.Println(err)
		fmt.Println("----------------------")
	}
	return parsedJSON
}
