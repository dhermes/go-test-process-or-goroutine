package winkin_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/dhermes/go-test-process-or-goroutine/winkin"
)

func TestSetValue_GetValue(t *testing.T) {
	pid := os.Getpid()
	ppid := os.Getppid()
	fmt.Printf("    nod_test: pid=%d, ppid=%d", pid, ppid)

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
