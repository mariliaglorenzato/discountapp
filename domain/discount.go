package domain

import (
	"time"
)

const eventsDiscountsFileName = "events_discounts"

type Discount struct {
	TotalDiscount float64
	Product       *Product
}

func NewDiscount(client *Client) *Discount {
	eventsConfig, err := LoadConfiguration()
	if err != nil {
		panic(err.Error())
	}

	eventDiscount := getEventDiscount(eventsConfig.events)
	ageDiscount := getAgeDiscount(client.getClientAge())
	totalDiscount := GetTotalDiscount(eventDiscount, ageDiscount)

	return &Discount{
		TotalDiscount: totalDiscount,
	}
}

func (d *Discount) GetPriceWithDiscount() uint64 {
	if d.Product != nil {
		floatPrice := float64(d.Product.Price / 100)
		priceInFloat := floatPrice - (floatPrice * d.TotalDiscount)
		return uint64(priceInFloat * 100)
	}

	return 0
}

func GetTotalDiscount(eventDiscount float64, ageDiscount float64) float64 {
	return eventDiscount + ageDiscount
}

func getEventDiscount(events []string) float64 {
	for _, event := range events {
		parsedDate, err := time.Parse("02/01", event)
		if err != nil {
			now := time.Now()
			if now.Format("02/01") == parsedDate.Format("02/01") {
				return float64(0.1)
			}

		}
	}
	return 0.0
}

func getAgeDiscount(age uint) float64 {
	return (float64(0.1) * float64(age))
}
