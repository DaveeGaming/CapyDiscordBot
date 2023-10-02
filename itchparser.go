package main

import (
	"log"
	"time"
)

var (
    syncTicker *time.Ticker
    lastSynced time.Time
)

func setTimer(timer time.Duration) {
    syncTicker = time.NewTicker(timer)
    lastSynced = time.Now()
}

func StartScraper() {
    for range syncTicker.C {
        log.Println("Timer fired")
        lastSynced = time.Now()
    }
}

func LoadJams() {

}
