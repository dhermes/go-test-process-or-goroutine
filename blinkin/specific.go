package blinkin

import (
	"github.com/dhermes/go-test-process-or-goroutine/winkin"
)

// SetSpecificValue sets a shared value to 42.
func SetSpecificValue() {
	winkin.SetValue(42)
}
