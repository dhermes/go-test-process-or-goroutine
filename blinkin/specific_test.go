package blinkin_test

import (
	"testing"

	"github.com/dhermes/go-test-process-or-goroutine/blinkin"
	"github.com/dhermes/go-test-process-or-goroutine/winkin"
)

func TestSetSpecificValue(t *testing.T) {
	blinkin.SetSpecificValue()
	got := winkin.GetValue()
	if got != 42 {
		t.Errorf("GetValue() = %d; want 42", got)
	}
}
