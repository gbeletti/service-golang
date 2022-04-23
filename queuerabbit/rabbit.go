package queuerabbit

import (
	"context"
	"log"
	"time"

	"github.com/gbeletti/rabbitmq"
)

var rabbit rabbitmq.RabbitMQ

// Start starts the RabbitMQ connection
func Start(ctx context.Context) {
	rabbit = rabbitmq.NewRabbitMQ()
	go setupRabbit(ctx)
}

func setupRabbit(ctx context.Context) {
	configConn := rabbitmq.ConfigConnection{}
	for {
		notifyClose, err := rabbit.Connect(configConn)
		if err != nil {
			log.Printf("error connecting to rabbitmq: %s\n", err)
			time.Sleep(time.Second * 5)
			continue
		}
		select {
		case <-notifyClose:
			continue
		case <-ctx.Done():
			return
		}
	}
}
