package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("Current local time: %s\n", now)

	jakarta, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Printf("Error loading Jakarta timezone: %v\n", err)
		return
	}

	london, err := time.LoadLocation("Europe/London")
	if err != nil {
		fmt.Printf("Error loading London timezone: %v\n", err)
		return
	}

	newYork, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Printf("Error loading New York timezone: %v\n", err)
		return
	}

	fmt.Println("\nSame moment in different timezones:")
	fmt.Printf("Jakarta: %s\n", now.In(jakarta))
	fmt.Printf("London: %s\n", now.In(london))
	fmt.Printf("New York: %s\n", now.In(newYork))

	meetingTime := time.Date(2026, 2, 16, 9, 0, 0, 0, jakarta)
	fmt.Printf("\nMeeting scheduled for: %s\n", meetingTime)
	fmt.Printf("That's %s in Jakartan\n", meetingTime.In(jakarta))
	fmt.Printf("That's %s in London\n", meetingTime.In(london))
	fmt.Printf("That's %s in New York\n", meetingTime.In(newYork))
	fmt.Printf("That's %s in UTC\n", meetingTime.UTC())

	timeStr := "2026-02-16T07:30:00-07:00"
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		fmt.Printf("Error parsing time: %v\n", err)
		return
	}
	fmt.Printf("\nParsed time: %s\n", parsedTime)
	fmt.Printf("In New York: %s\n", parsedTime.In(newYork))
	fmt.Printf("In Jakarta: %s\n", parsedTime.In(jakarta))

}
