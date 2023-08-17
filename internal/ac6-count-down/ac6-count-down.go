package ac6countdown

import (
	"fmt"
	"log"
	"math"
	"time"
)

func GetRemainingTime() string {
	location, err := time.LoadLocation("Europe/Moscow")

	if err != nil {
		log.Fatal(err)
	}

	releaseDate := time.Date(2023, time.August, 25, 0, 0, 0, 0, location)
	currentTime := time.Now()

	delta := releaseDate.Sub(currentTime)

	days, hours := convertDelta(delta)

	return fmt.Sprintf(
		"Main system will engage combat mode in %v days, %v hours",
		days,
		hours,
	)
}

func convertDelta(delta time.Duration) (float64, float64) {
	const hoursInDay = 24

	var days float64
	hours := delta.Abs().Hours()

	if hours > hoursInDay {
		days = math.Floor(hours / hoursInDay)
		hours -= days * hoursInDay
	}

	return days, math.Floor(hours)
}
