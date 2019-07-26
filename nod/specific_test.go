package nod_test

import (
	"testing"

	"github.com/dhermes/go-test-process-or-goroutine/nod"
	"github.com/dhermes/go-test-process-or-goroutine/winkin"
)

func TestSetSpecificValue(t *testing.T) {
	nod.SetSpecificValue()
	got := winkin.GetValue()
	if got != 1337 {
		t.Errorf("GetValue() = %d; want 1337", got)
	}
}
