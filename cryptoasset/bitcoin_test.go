package cryptoasset_test

import (
	"flag"
	"strings"
	"testing"

	"github.com/gbeletti/service-golang/cryptoasset"
)

var (
	update   = flag.Bool("update", false, "update golden file")
	testCase = flag.String("case", "", "run specific test case")
)

func TestCalculateBitcoinVariation(t *testing.T) {
	tcases := []struct {
		name       string
		startDate  string
		endDate    string
		resultFile string
	}{
		{
			name:       "01 - full month with all the data",
			startDate:  "2018-01-01",
			endDate:    "2018-01-31",
			resultFile: "./testdata/bitcoin/01/result.json",
		},
		{
			name:       "02 - data does not have the last days",
			startDate:  "2018-11-01",
			endDate:    "2018-11-25",
			resultFile: "./testdata/bitcoin/02/result.json",
		},
		{
			name:       "03 - only 2 days",
			startDate:  "2018-11-12",
			endDate:    "2018-11-13",
			resultFile: "./testdata/bitcoin/03/result.json",
		},
	}
	mock := newMockDB(t, "./testdata/database/bitcoin.json")
	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			testCalculateBitcoinVariation(t, tcase.name, tcase.startDate, tcase.endDate, tcase.resultFile, mock)
		})
	}

}

func testCalculateBitcoinVariation(t *testing.T, name, stardDate, endDate, resultFile string, mock mockDB) {
	if testCase != nil && len(*testCase) > 0 {
		if !strings.Contains(resultFile, *testCase) {
			t.Skipf("skipping test case [%s]", name)
		}
	}
	start, end, err := cryptoasset.ValidateStartEndDates(stardDate, endDate)
	if err != nil {
		t.Fatalf("couldnt validate dates start [%s] end [%s]. error [%s]", stardDate, endDate, err)
	}
	btVariation := cryptoasset.CalculateBitcoinVariation(start, end, mock)
	if *update {
		writeJSONFile(t, resultFile, btVariation, true)
		return
	}
	var expected []cryptoasset.Variation
	readJSONFile(t, resultFile, &expected)
	validateBitcoinVariation(t, name, expected, btVariation)
}

func validateBitcoinVariation(t *testing.T, name string, expected, got []cryptoasset.Variation) {
	if len(expected) != len(got) {
		t.Errorf("test [%s] expected [%d] but got [%d]", name, len(expected), len(got))
		return
	}
	for i := range expected {
		if expected[i].Date != got[i].Date {
			t.Errorf("test [%s] expected date [%s] but got [%s]", name, expected[i].Date, got[i].Date)
		}
		if expected[i].Variation != got[i].Variation {
			t.Errorf("test [%s] expected variation [%f] but got [%f]", name, expected[i].Variation, got[i].Variation)
		}
	}
}
