package cryptoasset

import (
	"context"

	"github.com/gbeletti/service-golang/dbmongo"
	"go.mongodb.org/mongo-driver/bson"
)

// BitcoinPriceGetter is the interface that wraps the GetBitcoinPrice method.
type BitcoinPriceGetter interface {
	GetBitcoinPrice(date string) (btprice BitcoinPrice, err error)
}

type btcDB struct{}

// GetBitcoinPrice returns the bitcoin price for the given date
func (btc btcDB) GetBitcoinPrice(date string) (btprice BitcoinPrice, err error) {
	return getBitcoinPrice(date)
}

func getBitcoinPrice(date string) (btprice BitcoinPrice, err error) {
	client, err := dbmongo.GetClient()
	if err != nil {
		return
	}
	err = client.Database("data").Collection("bitcoin").FindOne(context.Background(), bson.M{"date": date}).Decode(&btprice)
	return
}
