package main

import (
	"log"
	"time"
)

var (
    syncTicker *time.Ticker
    lastSynced time.Time
    syncTime time.Duration = time.Second * 10 //temporary
    scraping bool
)
//TODO: this will come from the config, the duration for the clock reset
func startTimer(){scraping = true; syncTicker.Reset(syncTime)}
func stopTimer() {scraping = false; syncTicker.Stop()}

// Initialize the scrape timer with the parameter
func init() {
    syncTicker = time.NewTicker(syncTime)
    lastSynced = time.Now()
    scraping = true
    })
}

// Start the scraper loop, that will call colly to scrape and store the necessary data
func StartScraper() {
    for range syncTicker.C {
        if scraping { 
            scrapeItch() 
        } else {
            return
        }
    }
}

func scrapeItch() {
    lastSynced = time.Now()
    log.Println("Scraping website")
}
