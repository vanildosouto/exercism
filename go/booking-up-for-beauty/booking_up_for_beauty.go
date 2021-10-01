package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	timeParsed, _ := time.Parse("1/2/2006 15:04:05", date)

	return timeParsed
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	timeParsed, _ := time.Parse("January 2, 2006 15:04:05", date)

	return timeParsed.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	timeParsed, _ := time.Parse("15:04:05", date[len(date)-8:])
	startAfternoon, _ := time.Parse("15:04:05", "12:00:00")
	endAfternoon, _ := time.Parse("15:04:05", "18:00:00")

	return timeParsed.After(startAfternoon) && timeParsed.Before(endAfternoon)
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	timeParsed := Schedule(date)

	return timeParsed.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	t := time.Date(2021, time.September, 15, 0, 0, 0, 0, time.UTC)

	return t
}
