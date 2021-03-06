package main

import (
	"context"
	"log"
	"time"

	"github.com/gbeletti/service-golang/dbmongo"
	"github.com/gbeletti/service-golang/httpserver"
	"github.com/gbeletti/service-golang/queuerabbit"
	"github.com/gbeletti/service-golang/servicemanager"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}
}

func main() {
	_, cancel := start()
	defer shutdown(cancel)
	servicemanager.WaitShutdown()
}

func start() (ctx context.Context, cancel context.CancelFunc) {
	// This is the main context for the service. When it is canceled it means the service is going down.
	// All the tasks must be canceled
	ctx, cancel = context.WithCancel(context.Background())
	if err := dbmongo.Start(); err != nil {
		log.Printf("couldnt start database error [%s]\n", err)
	}
	queuerabbit.Start(ctx)
	httpserver.Start()
	return
}

func shutdown(cancel context.CancelFunc) {
	cancel()
	ctx, cancelTimeout := context.WithTimeout(context.Background(), time.Second*30)
	defer cancelTimeout()
	doneHTTP := httpserver.Shutdown(ctx)
	doneRabbit := queuerabbit.Shutdown(ctx)
	err := servicemanager.WaitUntilIsDoneOrCanceled(ctx, doneHTTP, doneRabbit)
	if err != nil {
		log.Printf("service stopped by timeout %s\n", err)
	}
	if err := dbmongo.Close(context.Background()); err != nil {
		log.Printf("couldnt close mongo connection error [%s]\n", err)
	}
	time.Sleep(time.Millisecond * 200)
	log.Println("bye bye")
}
