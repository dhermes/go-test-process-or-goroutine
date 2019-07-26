package nod_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/dhermes/go-test-process-or-goroutine/nod"
	"github.com/dhermes/go-test-process-or-goroutine/winkin"
)

func TestSetSpecificValue(t *testing.T) {
	pid := os.Getpid()
	ppid := os.Getppid()
	fmt.Printf(" winkin_test: pid=%d, ppid=%d", pid, ppid)

	nod.SetSpecificValue()
	got := winkin.GetValue()
	if got != 1337 {
		t.Errorf("GetValue() = %d; want 1337", got)
	}
}
