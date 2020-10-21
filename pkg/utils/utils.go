package utils

import (
	"fmt"
	"math"
	"strings"
	"time"
)

//messages
var (
	//default display message when duration is 0
	displayMessage = "0 hour 0 minutes 0 seconds"
)

// HumanizeDuration humanizes time.Duration output to a meaningful value,
// golang's default ``time.Duration`` output is badly formatted and unreadable.
func HumanizeDuration(duration time.Duration) string {
	if duration == 0 {
		return displayMessage
	}
	days := int64(duration.Hours() / 24)
	hours := int64(math.Mod(duration.Hours(), 24))
	minutes := int64(math.Mod(duration.Minutes(), 60))
	seconds := int64(math.Mod(duration.Seconds(), 60))

	chunks := []struct {
		singularName string
		amount       int64
	}{
		{"day", days},
		{"hour", hours},
		{"minute", minutes},
		{"second", seconds},
	}

	parts := []string{}

	for _, chunk := range chunks {
		switch chunk.amount {
		case 0:
			continue
		case 1:
			parts = append(parts, fmt.Sprintf("%d %s", chunk.amount, chunk.singularName))
		default:
			parts = append(parts, fmt.Sprintf("%d %ss", chunk.amount, chunk.singularName))
		}
	}

	return strings.Join(parts, " ")
}