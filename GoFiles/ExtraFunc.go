package main

import "time"

// TimeFormat formats the time to UTC+1
func TimeFormat(date time.Time, format string) string {

	// Hour variable
	hour := 60 * time.Minute

	// Format to string in 'format' format
	formatted := date.UTC().Add(hour).Format(format)

	// Return the formatted string
	return formatted
}
