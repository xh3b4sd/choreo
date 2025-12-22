package ticker

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Framer_Ticker_Second_Round(t *testing.T) {
	var testCases = []struct {
		now string
		mul int
		res string
	}{
		// Case 000
		{
			now: "2024-01-15T12:34:56.400000000Z",
			mul: 1,
			res: "2024-01-15T12:34:56Z",
		},
		// Case 001
		{
			now: "2024-01-15T12:34:56.500000000Z",
			mul: 1,
			res: "2024-01-15T12:34:57Z",
		},
		// Case 002
		{
			now: "2024-01-15T12:34:56.999999999Z",
			mul: 1,
			res: "2024-01-15T12:34:57Z",
		},
		// Case 003
		{
			now: "2024-01-15T12:34:56.000000000Z",
			mul: 1,
			res: "2024-01-15T12:34:56Z",
		},
		// Case 004
		{
			now: "2024-01-15T12:34:58.100000000Z",
			mul: 5,
			res: "2024-01-15T12:35:00Z",
		},
		// Case 005
		{
			now: "2024-01-15T12:34:57.400000000Z",
			mul: 5,
			res: "2024-01-15T12:34:55Z",
		},
		// Case 006
		{
			now: "2024-01-15T12:34:57.500000000Z",
			mul: 5,
			res: "2024-01-15T12:35:00Z",
		},
		// Case 007
		{
			now: "2024-01-15T12:34:02.000000000Z",
			mul: 10,
			res: "2024-01-15T12:34:00Z",
		},
		// Case 008
		{
			now: "2024-01-15T12:34:05.000000000Z",
			mul: 10,
			res: "2024-01-15T12:34:10Z",
		},
		// Case 009
		{
			now: "2024-01-15T23:59:59.999999999Z",
			mul: 10,
			res: "2024-01-16T00:00:00Z",
		},
		// Case 010
		{
			now: "2024-01-15T00:00:04.999999999Z",
			mul: 10,
			res: "2024-01-15T00:00:00Z",
		},
		// Case 011
		{
			now: "2024-01-15T00:00:05.000000000Z",
			mul: 10,
			res: "2024-01-15T00:00:10Z",
		},
		// Case 012
		{
			now: "2024-12-31T23:59:58.600000000Z",
			mul: 5,
			res: "2025-01-01T00:00:00Z",
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
				res = Second(now.UTC()).Round(tc.mul).Time().Format(time.RFC3339Nano)
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}

func Test_Framer_Ticker_Second_Tick(t *testing.T) {
	var testCases = []struct {
		tic string
		qnt int
		res string
	}{
		// Case 000
		{
			tic: "2024-01-15T12:34:56Z",
			qnt: 0,
			res: "2024-01-15T12:34:56Z",
		},
		// Case 001
		{
			tic: "2024-01-15T12:34:56Z",
			qnt: +1,
			res: "2024-01-15T12:34:57Z",
		},
		// Case 002
		{
			tic: "2024-01-15T12:34:56Z",
			qnt: -1,
			res: "2024-01-15T12:34:55Z",
		},
		// Case 003
		{
			tic: "2024-01-15T12:34:56Z",
			qnt: +60,
			res: "2024-01-15T12:35:56Z",
		},
		// Case 004
		{
			tic: "2024-01-15T12:34:56Z",
			qnt: -60,
			res: "2024-01-15T12:33:56Z",
		},
		// Case 005
		{
			tic: "2024-12-31T23:59:59Z",
			qnt: +2,
			res: "2025-01-01T00:00:01Z",
		},
		// Case 006
		{
			tic: "2024-01-01T00:00:01Z",
			qnt: -2,
			res: "2023-12-31T23:59:59Z",
		},
		// Case 007
		{
			tic: "2024-01-15T12:00:00Z",
			qnt: +86400,
			res: "2024-01-16T12:00:00Z",
		},
		// Case 008
		{
			tic: "2024-01-15T12:00:00Z",
			qnt: -86400,
			res: "2024-01-14T12:00:00Z",
		},
		// Case 009
		{
			tic: "2024-01-15T12:34:56Z",
			qnt: +100000,
			res: "2024-01-16T16:21:36Z",
		},
		// Case 010
		{
			tic: "2024-01-15T12:34:56Z",
			qnt: -100000,
			res: "2024-01-14T08:48:16Z",
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
				res = Second(tic.UTC()).Tick(tc.qnt).Time().Format(time.RFC3339Nano)
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}
