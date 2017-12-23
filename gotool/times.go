package gotool

import (
	"fmt"
	"time"
)

const (
	FORMAT_Y_M_D_H_M_S = "%d-%02d-%02dT%02d-%02d-%02d"
)

// GetFormattedTimeInfo returns formatted time info.
func GetFormattedTimeInfo() string {
	t := time.Now()
	return fmt.Sprintf(FORMAT_Y_M_D_H_M_S, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
