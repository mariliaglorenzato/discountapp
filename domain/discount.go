package domain

type Discount struct {
	TotalDiscount float64
}

func GetPriceWithDiscount(price uint64, age uint, events []Event) (uint64, float64) {
	if events == nil {
		events = GetEventConfigInstance().Events
	}

	totalDiscount := GetTotalDiscount(age, events)

	floatPrice := float64(price / 100)
	floatPriceWithDiscount := floatPrice - (floatPrice * totalDiscount / 100)
	return uint64(floatPriceWithDiscount * 100), totalDiscount
}

func GetTotalDiscount(age uint, events []Event) float64 {
	ageDiscount := getAgeDiscount(age)

	eventDiscount := getEventDiscount(events)

	return eventDiscount + ageDiscount
}

func getAgeDiscount(age uint) float64 {
	return float64(age) * float64(0.1)
}
