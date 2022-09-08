package domain

import (
	"encoding/json"
	"os"
)

const eventsDiscountsFileName = "events_discounts.json"

type EventsConfig struct {
	Events []string `json:events`
}

func LoadConfiguration() (EventsConfig, error) {
	var eventsConfig EventsConfig
	configFile, err := os.Open(eventsDiscountsFileName)
	if err != nil {
		return eventsConfig, err
	}

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&eventsConfig.Events)

	return eventsConfig, nil
}
