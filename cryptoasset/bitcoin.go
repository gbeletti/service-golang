package cryptoasset

import (
	"errors"
	"time"
)

const (
	dateLayout string = "2006-01-02"
)

var (
	// ErrStartDateAfterEndDate is the error returned when the start date is after the end date
	ErrStartDateAfterEndDate = errors.New("start date after end date")
)

// BitcoinPrice represents the price of a bitcoin on a given date.
type BitcoinPrice struct {
	Date  string  `json:"date"`
	Price float64 `json:"price"`
}

// Variation represents a single day variation of an asset
type Variation struct {
	Date      string  `json:"date"`
	Variation float64 `json:"variation"`
}

// GetBitcoinVariation returns the bitcoin variation on a given period.
// The expected format is YYYY-MM-DD.
func GetBitcoinVariation(startDate, endDate string) (btcVars []Variation, err error) {
	start, end, err := validateStartEndDates(startDate, endDate)
	if err != nil {
		return
	}
	btcVars = calculateBitcoinVariation(start, end, btcDB{})

	return
}

func calculateBitcoinVariation(start, end time.Time, btcGetter BitcoinPriceGetter) (btcVars []Variation) {
	end = end.AddDate(0, 0, 1)
	var firstDone bool
	var btcPricePrevius BitcoinPrice
	var err error
	for start.Before(end) {
		var btcPrice BitcoinPrice
		btcPrice, err = btcGetter.GetBitcoinPrice(start.Format(dateLayout))
		if err != nil {
			btcVars = append(btcVars, Variation{
				Date:      start.Format(dateLayout),
				Variation: 0.0,
			})
			start = start.AddDate(0, 0, 1)
			continue
		}
		if !firstDone {
			btcPricePrevius = btcPrice
			firstDone = true
		}
		variat := btcPrice.Price - btcPricePrevius.Price
		btcVars = append(btcVars, Variation{
			Date:      start.Format(dateLayout),
			Variation: variat,
		})
		start = start.AddDate(0, 0, 1)
		btcPricePrevius = btcPrice
	}
	return
}

func validateStartEndDates(startDate, endDate string) (start, end time.Time, err error) {
	start, err = time.Parse(dateLayout, startDate)
	if err != nil {
		return
	}
	end, err = time.Parse(dateLayout, endDate)
	if err != nil {
		return
	}
	if start.After(end) {
		err = ErrStartDateAfterEndDate
		return
	}
	return
}
