package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const times_of_monitoring = 5
const delay_between_monitoring = 5 // Seconds

func main() {
	introduction()
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

type TSite struct {
	url         string
	environment string
	project     string
}

func initMonitore() {
	fmt.Println("Starting monitoring...")
	sites_to_monitore := []TSite{
		{
			url:         "https://staging-administrativo.mova.vc",
			environment: "staging",
			project:     "administrativo",
		},
		{
			url:         "https://administrativo.mova.vc",
			environment: "production",
			project:     "operacional",
		},
		{
			url:         "https://staging-operacional.mova.vc",
			environment: "staging",
			project:     "operacional",
		},
		{
			url:         "https://operacional.mova.vc",
			environment: "production",
			project:     "operacional",
		},
		{
			url:         "https://mova.vc",
			environment: "production",
			project:     "mova",
		},
	}

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

func testSiteStatus(site TSite) {
	response, _ := http.Get(site.url)
	if response.StatusCode == 200 {
		fmt.Println("SITE:", site.url)
		fmt.Println("PROJECT:", site.project)
		fmt.Println("STATUS", response.StatusCode)
		fmt.Println("ENVIRONMENT:", site.environment)
		fmt.Println("Site is OK!")
	} else {
		fmt.Println("SITE:", site.url)
		fmt.Println("PROJECT:", site.project)
		fmt.Println("STATUS:", response.StatusCode)
		fmt.Println("ENVIRONMENT:", site.environment)
		fmt.Println("Site", site, "have a problem :(")
	}
}
