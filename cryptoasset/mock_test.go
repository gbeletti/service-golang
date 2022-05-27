package cryptoasset_test

import (
	"errors"
	"testing"

	"github.com/gbeletti/service-golang/cryptoasset"
)

type mockDB struct {
	data map[string]cryptoasset.BitcoinPrice
}

func newMockDB(t *testing.T, fileDB string) (mock mockDB) {
	var prices []cryptoasset.BitcoinPrice
	readJSONFile(t, fileDB, &prices)
	mock = mockDB{
		data: make(map[string]cryptoasset.BitcoinPrice),
	}
	for _, price := range prices {
		mock.data[price.Date] = price
	}
	return
}

func (m mockDB) GetBitcoinPrice(date string) (btprice cryptoasset.BitcoinPrice, err error) {
	btprice, ok := m.data[date]
	if !ok {
		err = errors.New("not found")
	}
	return
}
