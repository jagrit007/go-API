package utils

import (
	"fmt"
	"log"
	"time"
)

func ParseDateTime(dt string) (time.Time, error) {
	// Try parsing with time
	parsedDT, err := time.Parse("02/01/2006 15:04", dt)
	if err == nil {
		return parsedDT, nil
	}

	// If it doesn't contain time, assume 00:00 and try again
	log.Printf("Time not found in the input. Assuming 00:00 as the default time.")
	dt = dt + " 00:00" // Append 00:00 if no time was provided

	// Try parsing the date with default time
	parsedDT, err = time.Parse("02/01/2006 15:04", dt)
	if err != nil {
		log.Printf("Failed to parse time in the desired format. Please check the value.")
		return time.Time{}, fmt.Errorf("failed to parse time: %v. Please check the value", err)
	}

	return parsedDT, nil
}
