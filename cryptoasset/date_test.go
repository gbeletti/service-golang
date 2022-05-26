package cryptoasset_test

import (
	"testing"
	"time"

	"github.com/gbeletti/service-golang/cryptoasset"
)

type dateInput struct {
	startDate string
	endDate   string
}

type dateResult struct {
	start    time.Time
	end      time.Time
	errIsNil bool
}

func TestValidateStartEndDates(t *testing.T) {
	tcases := []struct {
		name   string
		input  dateInput
		result dateResult
	}{
		{
			name: "valid dates",
			input: dateInput{
				startDate: "2018-01-01",
				endDate:   "2018-01-31",
			},
			result: dateResult{
				start:    time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				end:      time.Date(2018, 1, 31, 0, 0, 0, 0, time.UTC),
				errIsNil: true,
			},
		},
		{
			name: "same dates",
			input: dateInput{
				startDate: "2018-01-01",
				endDate:   "2018-01-01",
			},
			result: dateResult{
				start:    time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				end:      time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
				errIsNil: true,
			},
		},
		{
			name: "invalid dates",
			input: dateInput{
				startDate: "2018-13-01",
				endDate:   "2018-01-02",
			},
			result: dateResult{
				start:    time.Time{},
				end:      time.Time{},
				errIsNil: false,
			},
		},
		{
			name: "invalid period",
			input: dateInput{
				startDate: "2018-12-01",
				endDate:   "2018-01-02",
			},
			result: dateResult{
				start:    time.Date(2018, 12, 1, 0, 0, 0, 0, time.UTC),
				end:      time.Date(2018, 1, 2, 0, 0, 0, 0, time.UTC),
				errIsNil: false,
			},
		},
	}
	for _, tcase := range tcases {
		t.Run(tcase.name, func(t *testing.T) {
			testValidateStartEndDates(t, tcase.input, tcase.result)
		})
	}
}

func testValidateStartEndDates(t *testing.T, input dateInput, expected dateResult) {
	start, end, err := cryptoasset.ValidateStartEndDates(input.startDate, input.endDate)
	if err != nil && expected.errIsNil {
		t.Errorf("error is not nil: [%s] and it should be nil", err)
		return
	}
	if err == nil && !expected.errIsNil {
		t.Errorf("error is nil, but expected error is not nil")
		return
	}
	if !start.Equal(expected.start) {
		t.Errorf("start date is not correct: expected [%s] got [%s]", expected.start, start)
	}
	if !end.Equal(expected.end) {
		t.Errorf("end date is not correct: expected [%s] got [%s]", expected.end, end)
	}
}
