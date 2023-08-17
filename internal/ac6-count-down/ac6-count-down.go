package ac6countdown

import (
	"fmt"
	"math"
	"time"
)

func GetRemainingTime() string {
	releaseDate := time.Date(2023, time.August, 25, 0, 0, 0, 0, time.Local)
	currentTime := time.Now()

	delta := releaseDate.Sub(currentTime)

	days, hours := convertDelta(delta)

	return fmt.Sprintf("Main system will engage combat mode in %v days, %v hours", days, hours)
}

func convertDelta(delta time.Duration) (float64, float64) {
	const hoursInDay = 24

	var days float64
	hours := delta.Abs().Hours()

	if hours > hoursInDay {
		days = math.Floor(hours / hoursInDay)
		hours = hours - days*hoursInDay
	}

	return days, math.Floor(hours)
}
