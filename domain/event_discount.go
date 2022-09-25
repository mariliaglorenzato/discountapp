package domain

import (
	"time"

	"discountapp/utils"
)

type EventDiscount struct{}

func getEventDiscount(events []Event) float64 {
	for _, event := range events {
		parsedDate, _ := time.Parse(utils.DateFormatWithoutYear, event.Date)
		if dateMatchesActualTime(parsedDate) {
			return float64(0.1)
		}
	}
	return 0.0
}

func dateMatchesActualTime(dateToCompare time.Time) bool {
	now := time.Now()
	return now.Format(utils.DateFormatWithoutYear) == dateToCompare.Format(utils.DateFormatWithoutYear)
}
