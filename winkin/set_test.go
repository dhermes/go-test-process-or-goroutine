package winkin_test

import (
	"testing"

	"github.com/dhermes/go-test-process-or-goroutine/winkin"
)

func TestSetValue_GetValue(t *testing.T) {
	got := winkin.GetValue()
	if got != -1 {
		t.Errorf("GetValue() = %d; want -1", got)
	}
	winkin.SetValue(18)
	got = winkin.GetValue()
	if got != 18 {
		t.Errorf("GetValue() = %d; want 18", got)
	}
}
