package blinkin_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/dhermes/go-test-process-or-goroutine/blinkin"
	"github.com/dhermes/go-test-process-or-goroutine/winkin"
)

func TestSetSpecificValue(t *testing.T) {
	pid := os.Getpid()
	ppid := os.Getppid()
	fmt.Printf("blinkin_test: pid=%d, ppid=%d", pid, ppid)

	blinkin.SetSpecificValue()
	got := winkin.GetValue()
	if got != 42 {
		t.Errorf("GetValue() = %d; want 42", got)
	}
}
