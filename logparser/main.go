package main

import (
	"log"
	"os"

	"logparser/parser"

	"github.com/fsnotify/fsnotify"
)

func main() {
	log.Println("[main] Starting log parser...")

	logsDir := os.Getenv("LOGS_DIR")
	if logsDir == "" {
		logsDir = "/app/logs"
	}

	parsedDir := os.Getenv("PARSED_DIR")
	if parsedDir == "" {
		parsedDir = "/app/logs/parsed"
	}

	udpAddr := os.Getenv("UDP_ADDRESS")
	if udpAddr == "" {
		udpAddr = "logstash:5000"
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("[main] Failed to create watcher: %v", err)
	}
	defer watcher.Close()

	err = watcher.Add(logsDir)
	if err != nil {
		log.Fatalf("[main] Failed to watch directory %s: %v", logsDir, err)
	}

	log.Printf("[main] Watching directory: %s", logsDir)

	for {
		select {
		case event := <-watcher.Events:
			log.Printf("[main] Event detected: %s", event)
			if event.Op&fsnotify.Create == fsnotify.Create {
				go func(path string) {
					log.Printf("[main] Detected new file: %s", path)
					err := parser.ProcessLogFile(path, udpAddr, parsedDir)
					if err != nil {
						log.Printf("[main] Error processing file %s: %v", path, err)
					}
				}(event.Name)
			}
		case err := <-watcher.Errors:
			log.Printf("[main] Watcher error: %v", err)
		}
	}
}
