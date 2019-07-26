package nod

import (
	"github.com/dhermes/go-test-process-or-goroutine/winkin"
)

// SetSpecificValue sets a shared value to 1337.
func SetSpecificValue() {
	winkin.SetValue(1337)
}
