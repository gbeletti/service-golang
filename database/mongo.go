package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	opt    *options.ClientOptions
)

// Start should be called on the service start up. Makes the first connection to the database
func Start(opts ...*options.ClientOptions) (err error) {
	if len(opts) == 0 {
		opt = options.Client().ApplyURI(getURI())
	} else {
		opt = opts[0]
	}
	err = connectDB()
	return
}

// GetClient returns a MongoDB client
func GetClient() (*mongo.Client, error) {
	if client != nil {
		return client, nil
	}
	err := connectDB()
	return client, err
}

// Close closes all the connections to the database
func Close(ctx context.Context) (err error) {
	if client == nil {
		return
	}
	err = client.Disconnect(ctx)
	return
}

func connectDB() error {
	cli, err := mongo.NewClient(opt)
	if err != nil {
		return err
	}
	err = cli.Connect(context.Background())
	if err != nil {
		return err
	}
	client = cli
	return err
}

func getURI() (uri string) {
	uri = os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	return
}
