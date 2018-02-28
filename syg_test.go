package syg_test

import (
	"os"
	"runtime"
	"syscall"
	"testing"
	"time"

	"github.com/nasa9084/syg"
)

type checkCalled struct {
	b bool
}

func (chk *checkCalled) callback(os.Signal) {
	chk.b = true
}

func TestSyg(t *testing.T) {
	chk := &checkCalled{}
	cancel := syg.Listen(chk.callback, os.Interrupt)
	defer cancel()
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	time.Sleep(1 * time.Millisecond) // for-select looptime
	if !chk.b {
		t.Error("callback should be called")
		return
	}
}

func TestCancel(t *testing.T) {
	baseNumGoroutine := runtime.NumGoroutine()
	cancel := syg.Listen(func(os.Signal) {}, os.Interrupt)
	if runtime.NumGoroutine() != baseNumGoroutine+1 {
		t.Errorf("listen goroutine should be running: %d != %d", runtime.NumGoroutine(), baseNumGoroutine+1)
		return
	}
	cancel()
	time.Sleep(1 * time.Millisecond) // for-select looptime
	if runtime.NumGoroutine() != baseNumGoroutine {
		t.Errorf("listen goroutine should be canceled: %d != %d", runtime.NumGoroutine(), baseNumGoroutine)
		return
	}
}
