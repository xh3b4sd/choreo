package ticker

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Framer_Ticker_Hour_Round(t *testing.T) {
	var testCases = []struct {
		now string
		mul int
		res string
	}{
		// Case 000
		{
			now: "2024-01-15T12:34:56Z",
			mul: 1,
			res: "2024-01-15T13:00:00Z",
		},
		// Case 001
		{
			now: "2024-01-15T12:00:00Z",
			mul: 1,
			res: "2024-01-15T12:00:00Z",
		},
		// Case 002
		{
			now: "2024-01-15T12:00:00.000000001Z",
			mul: 1,
			res: "2024-01-15T12:00:00Z",
		},
		// Case 003
		{
			now: "2024-01-15T12:59:59.999999999Z",
			mul: 1,
			res: "2024-01-15T13:00:00Z",
		},
		// Case 004
		{
			now: "2024-01-15T14:38:00Z",
			mul: 6,
			res: "2024-01-15T12:00:00Z",
		},
		// Case 005
		{
			now: "2024-01-15T18:00:00Z",
			mul: 6,
			res: "2024-01-15T18:00:00Z",
		},
		// Case 006
		{
			now: "2024-01-15T23:59:59Z",
			mul: 6,
			res: "2024-01-16T00:00:00Z",
		},
		// Case 007
		{
			now: "2024-01-15T14:38:00Z",
			mul: 4,
			res: "2024-01-15T16:00:00Z",
		},
		// Case 008
		{
			now: "2024-01-15T03:59:59Z",
			mul: 4,
			res: "2024-01-15T04:00:00Z",
		},
		// Case 009
		{
			now: "2024-01-15T13:00:01Z",
			mul: 12,
			res: "2024-01-15T12:00:00Z",
		},
		// Case 010
		{
			now: "2024-01-15T01:00:01Z",
			mul: 12,
			res: "2024-01-15T00:00:00Z",
		},
		// Case 011
		{
			now: "2024-01-15T18:59:59.999999999Z",
			mul: 6,
			res: "2024-01-15T18:00:00Z",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var now time.Time
			{
				now, err = time.Parse(time.RFC3339Nano, tc.now)
				if err != nil {
					t.Fatal(err)
				}
			}

			var res string
			{
				res = Hour(now.UTC()).Round(tc.mul).Time().Format(time.RFC3339Nano)
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}

func Test_Framer_Ticker_Hour_Tick(t *testing.T) {
	var testCases = []struct {
		tic string
		qnt int
		res string
	}{
		// Case 000
		{
			tic: "2024-01-15T12:00:00Z",
			qnt: 0,
			res: "2024-01-15T12:00:00Z",
		},
		// Case 001
		{
			tic: "2024-01-15T12:00:00Z",
			qnt: +1,
			res: "2024-01-15T13:00:00Z",
		},
		// Case 002
		{
			tic: "2024-01-15T12:00:00Z",
			qnt: -1,
			res: "2024-01-15T11:00:00Z",
		},
		// Case 003
		{
			tic: "2024-01-15T12:00:00Z",
			qnt: +24,
			res: "2024-01-16T12:00:00Z",
		},
		// Case 004
		{
			tic: "2024-01-15T12:00:00Z",
			qnt: -24,
			res: "2024-01-14T12:00:00Z",
		},
		// Case 005
		{
			tic: "2024-12-31T23:00:00Z",
			qnt: +2,
			res: "2025-01-01T01:00:00Z",
		},
		// Case 006
		{
			tic: "2024-01-01T00:00:00Z",
			qnt: -2,
			res: "2023-12-31T22:00:00Z",
		},
		// Case 007
		{
			tic: "2024-01-15T12:00:00Z",
			qnt: +1000,
			res: "2024-02-26T04:00:00Z",
		},
		// Case 008
		{
			tic: "2024-01-15T12:00:00Z",
			qnt: -1000,
			res: "2023-12-04T20:00:00Z",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var tic time.Time
			{
				tic, err = time.Parse(time.RFC3339Nano, tc.tic)
				if err != nil {
					t.Fatal(err)
				}
			}

			var res string
			{
				res = Hour(tic.UTC()).Tick(tc.qnt).Time().Format(time.RFC3339Nano)
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}
