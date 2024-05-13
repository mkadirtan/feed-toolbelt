package main

import (
	"log"
	"os"

	"github.com/mkadirtan/feed-toolbelt/pkg/inspect"
)

func main() {
	outputLogger := log.New(os.Stdout, "", 0)
	debugLogger := log.New(os.Stderr, "", 0)

	usageText := `Usage: feed-toolbelt find [hostname]
Example: feed-toolbelt find https://nooptoday.com`

	if len(os.Args) != 3 {
		debugLogger.Printf("Invalid usage.\n%s", usageText)
		os.Exit(0)
	}

	command := os.Args[1]
	if command != "find" {
		debugLogger.Printf("Invalid command: %s.\n%s", command, usageText)
	}

	targetHostname := os.Args[2]
	foundFeeds := inspect.InspectURL(targetHostname, true, true, false)
	for _, feed := range foundFeeds {
		outputLogger.Println(feed)
	}
}
