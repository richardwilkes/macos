package macos

import (
	"fmt"
	"sync/atomic"
)

var nextRefID int64

func NextRefKey() string {
	return fmt.Sprintf("macos:%d", atomic.AddInt64(&nextRefID, 1))
}
