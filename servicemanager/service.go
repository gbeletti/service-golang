package servicemanager

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// ErrServiceCanceled is the error returned when the service is canceled
var ErrServiceCanceled = errors.New("service canceled")

// WaitShutdown waits until is going to die
func WaitShutdown() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-sigc
	log.Printf("signal received [%v] canceling everything\n", s)
}

// WaitUntilIsDoneOrCanceled it waits until all the dones channels are closed or the context is canceled
func WaitUntilIsDoneOrCanceled(ctx context.Context, dones ...chan struct{}) (err error) {
	done := make(chan struct{})
	go func() {
		for _, d := range dones {
			<-d
		}
		close(done)
	}()
	select {
	case <-done:
		log.Println("all done")
	case <-ctx.Done():
		err = ErrServiceCanceled
		log.Println("canceled")
	}
	return
}
