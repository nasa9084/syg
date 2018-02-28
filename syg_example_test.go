package syg_test

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/nasa9084/syg"
)

func ExampleSyg() {
	cancel := syg.Listen(func(os.Signal) { fmt.Println("interrupted") }, os.Interrupt)
	defer cancel()

	// in real world, press Ctrl-C
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	time.Sleep(1 * time.Millisecond) // for-select looptime

	// Output: interrupted
}
