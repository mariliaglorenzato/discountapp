package domain

import (
	"time"

	"discountapp/utils"
)

type Discount struct {
	TotalDiscount float64
}

func GetPriceWithDiscount(price uint64, age uint64, events []Event) (uint64, float64) {
	if events == nil {
		events = GetEventConfigInstance().Events
	}

	totalDiscount := GetTotalDiscount(age, events)

	floatPrice := float64(price / 100)

	floatPriceWithDiscount := floatPrice - (floatPrice * totalDiscount)
	return uint64(floatPriceWithDiscount * 100), totalDiscount
}

func GetTotalDiscount(age uint64, events []Event) float64 {
	ageDiscount := getAgeDiscount(age)

	eventDiscount := GetEventDiscount(events)

	return eventDiscount + ageDiscount
}

func getAgeDiscount(age uint64) float64 {
	return float64(age) * float64(0.001)
}

func GetEventDiscount(events []Event) float64 {
	for _, event := range events {
		parsedDate, _ := time.Parse(utils.DateFormatWithoutYear, event.Date)
		if dateMatchesActualTime(parsedDate) {
			return float64(0.001)
		}
	}
	return 0.0
}

func dateMatchesActualTime(dateToCompare time.Time) bool {
	now := time.Now()
	return now.Format(utils.DateFormatWithoutYear) == dateToCompare.Format(utils.DateFormatWithoutYear)
}
