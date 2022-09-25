package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"discountapp/domain"
	"discountapp/utils"
)

func LoadFile[T any](fileName string) *T {
	fileContent, _ := os.Open(fileName)

	defer fileContent.Close()

	file, _ := ioutil.ReadAll(fileContent)

	var data T
	json.Unmarshal([]byte(file), &data)

	return &data
}

func ToJsonI(data any) string {
	bytes, _ := json.MarshalIndent(data, "", " ")
	return string(bytes)
}

func ToJson(data any) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

func WriteEventInFile() {
	now := time.Now()
	event := domain.Event{
		Name: "Test",
		Date: now.Format(utils.DateFormatWithoutYear),
	}

	eventsConfig := domain.EventsConfig{
		Events: []domain.Event{
			event,
		},
	}

	data, _ := json.MarshalIndent(eventsConfig, "", " ")

	err := ioutil.WriteFile("../../config/events_discounts.test.json", data, 0)
	if err != nil {
		log.Fatal(err)
	}
}

func ResetEventsFixture() {
	eventsConfig := domain.EventsConfig{
		Events: []domain.Event{
			{
				Name: "black friday",
				Date: "21/09",
			},
			{
				Name: "natal",
				Date: "25/12",
			},
		},
	}

	data, _ := json.MarshalIndent(eventsConfig, "", " ")

	err := ioutil.WriteFile("../../config/events_discounts.test.json", data, 0)
	if err != nil {
		log.Fatal(err)
	}
}
