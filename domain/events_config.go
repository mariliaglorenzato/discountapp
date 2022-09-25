package domain

import (
	"encoding/json"

	"discountapp/config"
)

type EventsConfig struct {
	Events []Event `json:"events"`
}

type Event struct {
	Name string `json:"event_name"`
	Date string `json:"event_date"`
}

var singleEventConfigInstance *EventsConfig

func GetEventConfigInstance() *EventsConfig {
	if singleEventConfigInstance == nil {
		eventsFromConfig := config.LoadEventConfigFromFile()

		err := json.Unmarshal(eventsFromConfig, &singleEventConfigInstance)
		if err != nil {
			panic("Events Config File could not be parsed: " + err.Error()) // todo: improve config load errors handling
		}
		return singleEventConfigInstance
	}

	return singleEventConfigInstance
}
