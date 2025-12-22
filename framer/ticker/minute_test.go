package ticker

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Framer_Ticker_Minute_Round(t *testing.T) {
	var testCases = []struct {
		now string
		mul int
		res string
	}{
		// Case 000
		{
			now: "2024-01-15T12:34:56Z",
			mul: 1,
			res: "2024-01-15T12:35:00Z",
		},
		// Case 001
		{
			now: "2024-01-15T12:34:00Z",
			mul: 1,
			res: "2024-01-15T12:34:00Z",
		},
		// Case 002
		{
			now: "2024-01-15T12:34:00.000000001Z",
			mul: 1,
			res: "2024-01-15T12:34:00Z",
		},
		// Case 003
		{
			now: "2024-01-15T12:34:30Z",
			mul: 1,
			res: "2024-01-15T12:35:00Z",
		},
		// Case 004
		{
			now: "2024-01-15T14:38:00Z",
			mul: 15,
			res: "2024-01-15T14:45:00Z",
		},
		// Case 005
		{
			now: "2024-01-15T18:06:00Z",
			mul: 15,
			res: "2024-01-15T18:00:00Z",
		},
		// Case 006
		{
			now: "2024-01-15T18:07:30Z",
			mul: 15,
			res: "2024-01-15T18:15:00Z",
		},
		// Case 007
		{
			now: "2024-01-15T12:04:00Z",
			mul: 10,
			res: "2024-01-15T12:00:00Z",
		},
		// Case 008
		{
			now: "2024-01-15T12:05:00Z",
			mul: 10,
			res: "2024-01-15T12:10:00Z",
		},
		// Case 009
		{
			now: "2024-01-15T12:16:00Z",
			mul: 30,
			res: "2024-01-15T12:30:00Z",
		},
		// Case 010
		{
			now: "2024-01-15T23:59:59.999999999Z",
			mul: 15,
			res: "2024-01-16T00:00:00Z",
		},
		// Case 011
		{
			now: "2024-01-15T00:07:29.999999999Z",
			mul: 15,
			res: "2024-01-15T00:00:00Z",
		},
		// Case 012
		{
			now: "2024-01-15T00:07:30Z",
			mul: 15,
			res: "2024-01-15T00:15:00Z",
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
				res = Minute(now.UTC()).Round(tc.mul).Time().Format(time.RFC3339Nano)
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}

func Test_Framer_Ticker_Minute_Tick(t *testing.T) {
	var testCases = []struct {
		tic string
		qnt int
		res string
	}{
		// Case 000
		{
			tic: "2024-01-15T12:34:00Z",
			qnt: 0,
			res: "2024-01-15T12:34:00Z",
		},
		// Case 001
		{
			tic: "2024-01-15T12:34:00Z",
			qnt: +1,
			res: "2024-01-15T12:35:00Z",
		},
		// Case 002
		{
			tic: "2024-01-15T12:34:00Z",
			qnt: -1,
			res: "2024-01-15T12:33:00Z",
		},
		// Case 003
		{
			tic: "2024-01-15T12:34:00Z",
			qnt: +60,
			res: "2024-01-15T13:34:00Z",
		},
		// Case 004
		{
			tic: "2024-01-15T12:34:00Z",
			qnt: -60,
			res: "2024-01-15T11:34:00Z",
		},
		// Case 005
		{
			tic: "2024-12-31T23:59:00Z",
			qnt: +2,
			res: "2025-01-01T00:01:00Z",
		},
		// Case 006
		{
			tic: "2024-01-01T00:01:00Z",
			qnt: -2,
			res: "2023-12-31T23:59:00Z",
		},
		// Case 007
		{
			tic: "2024-01-15T12:00:00Z",
			qnt: +1440,
			res: "2024-01-16T12:00:00Z",
		},
		// Case 008
		{
			tic: "2024-01-15T12:00:00Z",
			qnt: -1440,
			res: "2024-01-14T12:00:00Z",
		},
		// Case 009
		{
			tic: "2024-01-15T12:34:00Z",
			qnt: +10000,
			res: "2024-01-22T11:14:00Z",
		},
		// Case 010
		{
			tic: "2024-01-15T12:34:00Z",
			qnt: -10000,
			res: "2024-01-08T13:54:00Z",
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
				res = Minute(tic.UTC()).Tick(tc.qnt).Time().Format(time.RFC3339Nano)
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}
