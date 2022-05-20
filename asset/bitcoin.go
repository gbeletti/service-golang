package asset

// GetBitcoinVariation returns the bitcoin variation on a given period.
// The expected format is YYYY-MM-DD.
func GetBitcoinVariation(startDate, endDate string) (btcVars []Variation) {
	return
}

type bitcoinPrice struct {
	Date  string  `json:"date"`
	Price float64 `json:"price"`
}

// Variation represents a single day variation of an asset
type Variation struct {
	Date      string
	Variation float64
}
