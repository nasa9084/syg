package syg

import (
	"context"
	"os"
	"os/signal"
)

// Listen launches new goroutine which waits given signals and returns CancelFunc.
// When one of the signals is called, given callback function will be called.
func Listen(callback func(os.Signal), signals ...os.Signal) context.CancelFunc {
	ctx := context.Background()
	return ListenContext(ctx, callback, signals...)
}

// ListenContext launches new goroutine which waits given signals and returns CancelFunc for given context.
// When one of the signals is called, given callback function will be called.
func ListenContext(ctx context.Context, callback func(os.Signal), signals ...os.Signal) context.CancelFunc {
	ctx, cancel := context.WithCancel(ctx)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, signals...)
	go listen(ctx, sigCh, callback)
	return cancel
}

func listen(ctx context.Context, sigCh chan os.Signal, callback func(os.Signal)) {
	for {
		select {
		case sig := <-sigCh:
			callback(sig)
		case <-ctx.Done():
			return
		}
	}
}
