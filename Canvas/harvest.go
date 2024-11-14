package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/firefox"
)

func main() {
	// Start the geckodriver service for Firefox
	service, err := selenium.NewGeckoDriverService("./geckodriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer service.Stop()

	// Set up Firefox capabilities
	caps := selenium.Capabilities{}
	caps.AddFirefox(firefox.Capabilities{
		Args: []string{
			// "--headless", // Uncomment this if you want to run Firefox in headless mode
		},
	})

	// Connect to the Selenium WebDriver
	driver, err := selenium.NewRemote(caps, "http://127.0.0.1:4444/wd/hub")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Maximize the browser window
	err = driver.MaximizeWindow("")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Navigate to the desired URL
	err = driver.Get("https://ozark.instructure.com/")
	if err != nil {
		log.Fatal("Error:", err)
	}

	timeout := 10 * time.Second

	// Find the student button using XPATH
	xpath := "/html/body/div[1]/div/div/main/div/div/div[2]/div[1]/div/div/p/a"
	studentButton, err := waitForElementXPATH(driver, xpath, timeout)
	if err != nil {
		log.Fatal("Error", err)
	}
	studentButton.Click()

	// Find the Google login button using class name
	ClassName := "flexbox items--center AuthMethod--container"
	googleButton, err := waitForElementClassName(driver, ClassName, timeout)
	if err != nil {
		log.Fatal("Error", err)
	}
	googleButton.SendKeys(string('\ue006'))

	// Wait for a minute before exiting
	time.Sleep(time.Minute)
}

func waitForElementClassName(driver selenium.WebDriver, of string, timeout time.Duration) (selenium.WebElement, error) {
	startTime := time.Now()

	for {
		// Try to find the element by class name
		element, err := driver.FindElement(selenium.ByClassName, of)
		if err == nil {
			return element, nil
		}

		// If the element is not found and timeout is reached, return an error
		if time.Since(startTime) > timeout {
			return nil, fmt.Errorf("Error: Element not found after %v seconds", timeout)
		}

		// Sleep before trying again
		time.Sleep(500 * time.Millisecond)
	}
}

func waitForElementXPATH(driver selenium.WebDriver, of string, timeout time.Duration) (selenium.WebElement, error) {
	startTime := time.Now()

	for {
		// Try to find the element by XPath
		element, err := driver.FindElement(selenium.ByXPATH, of)
		if err == nil {
			return element, nil
		}

		// If the element is not found and timeout is reached, return an error
		if time.Since(startTime) > timeout {
			return nil, fmt.Errorf("Error: Element not found after %v seconds", timeout)
		}

		// Sleep before trying again
		time.Sleep(500 * time.Millisecond)
	}
}
