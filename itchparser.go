package main

import (
	"log"
	"time"

	"github.com/gocolly/colly"
)

type Jam struct {
    Joined, ImageLink, JamLink, Duration string
    StartsIn time.Duration
}

var (
    syncTicker *time.Ticker
    lastSynced time.Time
    syncTime time.Duration = time.Second * 10 //temporary

    scraper *colly.Collector
    scraping bool

    jamEntries map[string]*Jam 
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
        jamEntries = map[string]*Jam{}
    })

    scraper.OnHTML("div.jam", func(h *colly.HTMLElement) {
        //Check jam list, if we saved it already, just update time, else, create new jam
        jamName := h.ChildText(".primary_info")
        if val, ok := jamEntries[jamName]; ok {
            startTime, _ := time.Parse("2006-01-02 15:04:05",h.ChildAttr(".date_countdown", "title"))
            val.StartsIn = startTime.Sub(time.Now())
            val.Duration = h.ChildText(".date_duration")
            val.Joined = h.ChildText(".number")
        } else {
            //Create empty jam to store data in
            currentJam := Jam{}
            startTime, _ := time.Parse("2006-01-02 15:04:05",h.ChildAttr(".date_countdown", "title"))
            currentJam.JamLink = "https://itch.io" + h.ChildAttr("a", "href")
            currentJam.ImageLink = h.ChildAttr(".jam_cover", "data-background_image")
            currentJam.StartsIn = startTime.Sub(time.Now())
            currentJam.Duration = h.ChildText(".date_duration")
            currentJam.Joined = h.ChildText(".number")
            jamEntries[jamName] = &currentJam
        }
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
    scraper.Visit("https://itch.io/jams/upcoming/sort-date")
    scraping = false
}
