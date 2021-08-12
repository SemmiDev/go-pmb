package utils

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/constants"
	"log"
	"time"
)

type Time struct {
	time.Time
	Date     time.Time
	Now      time.Time
	Location *time.Location
}

func NewTime(timeConfig ...Time) *Time {
	location, err := time.LoadLocation(constants.Timezone)
	if err != nil {
		log.Fatal(err)
	}

	var newConfig Time

	if len(timeConfig) > 0 {
		newConfig = timeConfig[0]
		newConfig.Location = location
		return &newConfig
	}

	return &Time{
		Now:      time.Now(),
		Location: location,
	}
}

func (t *Time) DaysInMonth(i time.Time) int {
	return i.AddDate(0, 1, 0).Add(time.Nanosecond * -1).Day()
}

func (t *Time) StartDate() time.Time {
	return time.Date(t.Now.Year(), t.Now.Month(), 1, 0, 0, 0, 0, t.Location)
}

func (t *Time) EndDate() time.Time {
	return time.Date(t.Now.Year(), t.Now.Month(), t.DaysInMonth(t.StartDate()), 23, 59, 59, 59, t.Location)
}

func (t *Time) FormatDate() time.Time {
	return time.Date(t.Date.Year(), t.Date.Month(), t.Date.Day(), t.Date.Hour(), t.Date.Minute(), t.Date.Second(), t.Date.Nanosecond(), t.Location)
}

func (t *Time) FormatDateWithInput(i time.Time) time.Time {
	return time.Date(i.Year(), i.Month(), i.Day(), i.Hour(), i.Minute(), i.Second(), i.Nanosecond(), t.Location)
}
