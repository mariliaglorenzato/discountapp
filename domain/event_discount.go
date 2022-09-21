package domain

import (
	"time"

	"discountapp/utils"
)

type EventDiscount struct{}

func getEventDiscount(events []Event) float64 {
	for _, event := range events {
		parsedDate, err := time.Parse(utils.DateFormatWithoutYear, event.Date)
		if err != nil {
			if dateMatchesWithActualTime(parsedDate) {
				return float64(0.1)
			}
		}
	}
	return 0.0
}

func dateMatchesWithActualTime(dateToCompare time.Time) bool {
	now := time.Now()
	return now.Format(utils.DateFormatWithoutYear) == dateToCompare.Format(utils.DateFormatWithoutYear)
}
