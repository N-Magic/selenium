package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/firefox"
)

func main() {
	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewGeckoDriverService("./geckodriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}

	defer service.Stop()

	// configure the browser options
	caps := selenium.Capabilities{}
	caps.AddFirefox(firefox.Capabilities{Args: []string{
		//"--headless-new", // comment out this line for testing
	}})

	// create a new remote client with the specified options
	fmt.Println("Starting new remote")
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println("Made new remote")

	err = wd.MaximizeWindow("")
	errorCheck(err)

	err = wd.Get("https://10best.usatoday.com/awards/travel/best-theme-park-holiday-event-2024/an-old-time-christmas-at-silver-dollar-city-branson-missouri/")
	errorCheck(err)

	time.Sleep(time.Minute)

}
func errorCheck(err error) {
	if err != nil {
		log.Fatal("Error : ", err)
	}
}
