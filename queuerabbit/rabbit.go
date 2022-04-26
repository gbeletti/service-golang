package queuerabbit

import (
	"context"
	"log"
	"os"

	"github.com/gbeletti/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

var rabbit rabbitmq.RabbitMQ

// Start starts the RabbitMQ connection
func Start(ctx context.Context) {
	rabbit = rabbitmq.NewRabbitMQ()

	setupRabbit(ctx)
}

// Shutdown stops the RabbitMQ connection
func Shutdown(ctx context.Context) (done chan struct{}) {
	done = rabbit.Close(ctx)
	return
}

func setupRabbit(ctx context.Context) {
	var setup rabbitmq.Setup = func() {
		createQueues(rabbit)
		createConsumers(ctx, rabbit)
	}
	configConn := rabbitmq.ConfigConnection{
		URI:           loadURI(),
		PrefetchCount: 1,
	}
	rabbitmq.KeepConnectionAndSetup(ctx, rabbit, configConn, setup)
}

func createQueues(rabbit rabbitmq.QueueCreator) {
	config := rabbitmq.ConfigQueue{
		Name:       "test",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
	_, err := rabbit.CreateQueue(config)
	if err != nil {
		log.Printf("error creating queue: %s\n", err)
	}
}

func createConsumers(ctx context.Context, rabbit rabbitmq.Consumer) {
	config := rabbitmq.ConfigConsume{
		QueueName:         "test",
		Consumer:          "test",
		AutoAck:           false,
		Exclusive:         false,
		NoLocal:           false,
		NoWait:            false,
		Args:              nil,
		ExecuteConcurrent: true,
	}
	go func() {
		if err := rabbit.Consume(ctx, config, receiveMessage); err != nil {
			log.Printf("error consuming from queue: %s\n", err)
		}
	}()
}

func receiveMessage(d *amqp.Delivery) {
	defer func() {
		if err := d.Ack(false); err != nil {
			log.Printf("error acking message: %s\n", err)
		}
	}()
	log.Printf("received message: %s\n", d.Body)
}

// PublishTest publishes a test message to the RabbitMQ exchange
func PublishTest(ctx context.Context, msg string) {
	config := rabbitmq.ConfigPublish{
		Exchange:   "",
		RoutingKey: "test",
	}
	if err := rabbit.Publish(ctx, []byte(msg), config); err != nil {
		log.Printf("error publishing message: %s\n", err)
	}
}

func loadURI() (uri string) {
	uri = os.Getenv("RABBITMQ_URI")
	return
}
