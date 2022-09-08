package domain

import (
	"time"
)

const (
	dateFormatWithoutYear = "02/01"
)

type Discount struct {
	TotalDiscount float64
}

func NewDiscount(client *Client) *Discount {
	eventsConfig, err := LoadConfiguration()
	if err != nil {
		panic(err.Error())
	}

	eventDiscount := getEventDiscount(eventsConfig.Events)
	ageDiscount := getAgeDiscount(client.getClientAge())
	totalDiscount := GetTotalDiscount(eventDiscount, ageDiscount)

	return &Discount{
		TotalDiscount: totalDiscount,
	}
}

func GetTotalDiscount(eventDiscount float64, ageDiscount float64) float64 {
	return eventDiscount + ageDiscount
}

func (d *Discount) GetPriceWithDiscount(price uint64) uint64 {
	floatPrice := float64(price / 100)
	priceInFloat := floatPrice - (floatPrice * d.TotalDiscount)
	return uint64(priceInFloat * 100)
}

func getEventDiscount(events []string) float64 {
	for _, event := range events {
		parsedDate, err := time.Parse(dateFormatWithoutYear, event)
		if err != nil {
			now := time.Now()
			if now.Format(dateFormatWithoutYear) == parsedDate.Format(dateFormatWithoutYear) {
				return float64(0.1)
			}

		}
	}
	return 0.0
}

func getAgeDiscount(age uint) float64 {
	return (float64(0.1) * float64(age))
}
