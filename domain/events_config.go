package domain

import (
	"encoding/json"
	"os"
)

type EventsConfig struct {
	events []string `json:events`
}

func LoadConfiguration() (EventsConfig, error) {
	var eventsConfig EventsConfig
	configFile, err := os.Open(eventsDiscountsFileName)
	if err != nil {
		return eventsConfig, err
	}

	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&eventsConfig.events)

	return eventsConfig, nil
}
