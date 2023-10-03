package main

import (
	"log"
	"time"

	"github.com/gocolly/colly"
)

type Jam struct {
    ImageLink, JamLink, Name, StartsIn, Duration string
}

var (
    syncTicker *time.Ticker
    lastSynced time.Time
    syncTime time.Duration = time.Second * 10 //temporary

    scraper *colly.Collector
    scraping bool

    jamEntries []Jam
)
//TODO: this will come from the config, the duration for the clock reset
func startTimer(){scraping = true; syncTicker.Reset(syncTime)}
func stopTimer() {scraping = false; syncTicker.Stop()}

// Initialize the scrape timer with the parameter
func init() {
    syncTicker = time.NewTicker(syncTime)
    lastSynced = time.Now()
    scraping = true

    scraper = colly.NewCollector()

    //Reset content of the jameEntries
    scraper.OnRequest(func(r *colly.Request) {
        jamEntries = []Jam{}
    })

    scraper.OnHTML("div.padded_content", func(h *colly.HTMLElement) {
        currentJam := Jam{}
        currentJam.Name = h.ChildText(".primary_info")
        currentJam.JamLink = "https://itch.io" + h.ChildAttr(".primary_info", "href")
        currentJam.ImageLink = h.ChildAttr(".jam_cover", "data-background_image")
        currentJam.StartsIn = h.ChildText(".date_countdown")
        currentJam.Duration = h.ChildText(".date_duration")
        jamEntries = append(jamEntries, currentJam)
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
    scraper.Visit("https://itch.io/jams/upcoming")
    scraping = false
}
