package gotools

import (
	"fmt"
	"time"
)

func ISO8601(t *time.Time) string {
	var tz string
	z, off := t.Zone()
	if z == "UTC" {
		tz = "Z"
	} else {
		tz = fmt.Sprintf("%03d00", off/3600)
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02d%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(),
		t.Second(), tz)
}
