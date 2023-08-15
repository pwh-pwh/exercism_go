package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	/*
		Schedule("7/25/2019 13:45:00")
		// => 2019-07-25 13:45:00 +0000 UTC
	*/
	parse, _ := time.Parse("1/2/2006 15:04:05", date)
	return parse
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	/**
	HasPassed("July 25, 2019 13:45:00")
	// => true
	*/
	parse, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(parse)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	/**
	IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
	// => true
	*/
	parse, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return parse.Hour() >= 12 && parse.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	/**
	Description("7/25/2019 13:45:00")
	// => "You have an appointment on Thursday, July 25, 2019, at 13:45."
	*/
	parse, _ := time.Parse("1/2/2006 15:04:05", date)
	return "You have an appointment on " + parse.Weekday().String() + ", " + parse.Month().String() + " " + parse.Format("2") + ", " + parse.Format("2006") + ", at " + parse.Format("15:04") + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	/**
	In this exercise you'll be working on an appointment scheduler for a beauty salon that opened on September 15th in 2012.

	You have five tasks, which will all involve appointment dates.


		Implement the AnniversaryDate function that returns the anniversary date of the salon's opening for the current year in UTC.

	Assuming the current year is 2020:

	AnniversaryDate()

	// => 2020-09-15 00:00:00 +0000 UTC

	*/
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
