package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"
)
type Duration struct {
    time.Duration
}

type Config struct {
	SyncTimer Duration 
}

var config Config

//Implement duration marshaling
func (d *Duration) MarshalJSON() ([]byte, error) {
    return json.Marshal(d.Duration.String())
}

func (d *Duration) UnmarshalJSON(b []byte) (error) {
    var v interface{}
    if err := json.Unmarshal(b, &v); err != nil {
        return err
    }
    switch value := v.(type) {
    case float64:
        d.Duration = time.Duration(value)
        return nil
    case string:
        var err error
        d.Duration, err = time.ParseDuration(value)
        if err != nil {
            return err
        }
        return nil
    default:
        return errors.New("invalid duration")
    }
}



func init() {
    fileContent, err := os.ReadFile("config.json")
    if err != nil {
        if CreateConfig {
            log.Println("Creating default config file")
            //Create default config settings
            config = Config{}
            config.SyncTimer = Duration{time.Hour * 2}

            saveConfig()
        } else {
            log.Panicf("Unable to find config file: %v", err)
        }
    }
	json.Unmarshal(fileContent, &config)
}


func saveConfig() {
    log.Println("Saving config")
    data, err := json.MarshalIndent(&config, "", " ")
    if err != nil {
        log.Panicf("Unable to create json config data: %v", err)
    }
    os.WriteFile("config.json", data, 0644)
    if err != nil {
        log.Panicf("Unable to write json data to file: %v", err)
    }
}
