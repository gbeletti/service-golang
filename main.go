package main

import (
	"context"
	"log"
	"time"

	"github.com/gbeletti/service-golang/httpserver"
	"github.com/gbeletti/service-golang/queuerabbit"
	"github.com/gbeletti/service-golang/servicemanager"
)

func main() {
	_, cancel := start()
	defer shutdown(cancel)
	servicemanager.WaitShutdown()
}

func start() (ctx context.Context, cancel context.CancelFunc) {
	// This is the main context for the service. When it is canceled it means the service is going down.
	// All the tasks must be canceled
	ctx, cancel = context.WithCancel(context.Background())
	queuerabbit.Start(ctx)
	httpserver.Start()
	return
}

func shutdown(cancel context.CancelFunc) {
	cancel()
	ctx, cancelTimeout := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelTimeout()
	doneHTTP := httpserver.Shutdown(ctx)
	err := servicemanager.WaitUntilIsDoneOrCanceled(ctx, doneHTTP)
	if err != nil {
		log.Printf("service stopped by timeout %s\n", err)
	}
	log.Println("bye bye")
}
