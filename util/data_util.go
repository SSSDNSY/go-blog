package util

import (
	"strings"
	"time"
)

func PaserDateTime(d time.Time) string {
	return strings.Split(d.String(), ".")[0]
}
