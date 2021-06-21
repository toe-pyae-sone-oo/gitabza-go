package dateutil

import (
	"fmt"
	"time"
)

// current gets the current date same as time.Now()
// for testing
var current = time.Now

func GetCurrentDateInStr() string {
	year, month, day := current().Date()
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}
