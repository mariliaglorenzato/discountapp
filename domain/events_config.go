package domain

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
)

var eventsDiscountsFileName string = "events_discounts.json"

type EventsConfig struct {
	Events []Event `json:"events"`
}

type Event struct {
	Name string `json:"event_name"`
	Date string `json:"event_date"`
}

var singleEventConfigInstance *EventsConfig

func NewEventsConfig() (*EventsConfig, error) {
	eventsConfig, err := loadConfiguration()
	if err != nil {
		return nil, err
	}

	return &eventsConfig, nil
}

func loadConfiguration() (EventsConfig, error) {
	content, err := ioutil.ReadFile(filepath.Clean(getConfigJsonEntirePath()))
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		panic(err.Error())
	}

	var eventsConfig EventsConfig
	err = json.Unmarshal(content, &eventsConfig)
	if err != nil {
		log.Fatal("Error during parsing event config file: ", err)
	}

	return eventsConfig, nil
}

func GetEventConfigInstance() *EventsConfig {
	if singleEventConfigInstance == nil {
		singleEventConfigInstance, _ := NewEventsConfig()
		return singleEventConfigInstance
	}

	return singleEventConfigInstance
}

func getConfigJsonEntirePath() string {
	return filepath.Join(getBasePath(), "../", eventsDiscountsFileName)
}

func getBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}
