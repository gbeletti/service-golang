package cryptoasset

import (
	"errors"
	"time"
)

const (
	dateLayout string = "2006-01-02"
)

var (
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
	for start.Before(end) {
		btcVars = append(btcVars, Variation{
			Date:      start.Format(dateLayout),
			Variation: 0.0,
		})
		start = start.AddDate(0, 0, 1)
	}
	return
}

func validateStartEndDates(startDate, endDate string) (start, end time.Time, err error) {
	start, err = parseDate(startDate)
	if err != nil {
		return
	}
	end, err = parseDate(endDate)
	if err != nil {
		return
	}
	if start.After(end) {
		err = ErrStartDateAfterEndDate
		return
	}
	return
}

func parseDate(dt string) (date time.Time, err error) {
	date, err = time.Parse(dateLayout, dt)
	return
}
