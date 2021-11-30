package scheduleAway

import (
	"time"
)

var Schedule []time.Time

func AddScheduledAway(hoursLater int, schedule *[]time.Time) {
	timeToAdd := time.Now().Add(time.Hour * time.Duration(hoursLater))
	*schedule = append(*schedule, timeToAdd)
}
