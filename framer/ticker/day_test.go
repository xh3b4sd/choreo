package ticker

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Framer_Ticker_Day_Round(t *testing.T) {
	var testCases = []struct {
		now string
		mul int
		res string
	}{
		// Case 000
		{
			now: "2024-01-15T11:59:59Z",
			mul: 1,
			res: "2024-01-15T00:00:00Z",
		},
		// Case 001
		{
			now: "2024-01-15T12:00:00Z",
			mul: 1,
			res: "2024-01-16T00:00:00Z",
		},
		// Case 002
		{
			now: "2024-01-15T12:34:56Z",
			mul: 1,
			res: "2024-01-16T00:00:00Z",
		},
		// Case 003
		{
			now: "2024-01-15T00:00:00Z",
			mul: 1,
			res: "2024-01-15T00:00:00Z",
		},
		// Case 004
		{
			now: "2024-12-31T11:00:00Z",
			mul: 1,
			res: "2024-12-31T00:00:00Z",
		},
		// Case 005
		{
			now: "2024-12-31T12:00:00Z",
			mul: 1,
			res: "2025-01-01T00:00:00Z",
		},
		// Case 006
		{
			now: "2024-12-31T23:59:59.999999999Z",
			mul: 1,
			res: "2025-01-01T00:00:00Z",
		},
		// Case 007
		{
			now: "2024-01-01T12:34:56Z",
			mul: 7,
			res: "2024-01-01T00:00:00Z",
		},
		// Case 008
		{
			now: "2024-01-04T13:00:00Z",
			mul: 7,
			res: "2024-01-08T00:00:00Z",
		},
		// Case 009
		{
			now: "2024-01-11T11:00:00Z",
			mul: 7,
			res: "2024-01-08T00:00:00Z",
		},
		// Case 010
		{
			now: "2024-01-11T12:00:00Z",
			mul: 7,
			res: "2024-01-15T00:00:00Z",
		},
		// Case 011
		{
			now: "2024-12-31T12:34:56Z",
			mul: 7,
			res: "2024-12-29T00:00:00Z",
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
				res = Day(now.UTC()).Round(tc.mul).Time().Format(time.RFC3339Nano)
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}

func Test_Framer_Ticker_Day_Tick(t *testing.T) {
	var testCases = []struct {
		tic string
		qnt int
		res string
	}{
		// Case 000
		{
			tic: "2024-01-15T00:00:00Z",
			qnt: 0,
			res: "2024-01-15T00:00:00Z",
		},
		// Case 001
		{
			tic: "2024-01-15T00:00:00Z",
			qnt: +1,
			res: "2024-01-16T00:00:00Z",
		},
		// Case 002
		{
			tic: "2024-01-15T00:00:00Z",
			qnt: -1,
			res: "2024-01-14T00:00:00Z",
		},
		// Case 003
		{
			tic: "2024-01-31T00:00:00Z",
			qnt: +1,
			res: "2024-02-01T00:00:00Z",
		},
		// Case 004
		{
			tic: "2024-03-01T00:00:00Z",
			qnt: -1,
			res: "2024-02-29T00:00:00Z",
		},
		// Case 005
		{
			tic: "2024-12-31T00:00:00Z",
			qnt: +1,
			res: "2025-01-01T00:00:00Z",
		},
		// Case 006
		{
			tic: "2024-01-01T00:00:00Z",
			qnt: -1,
			res: "2023-12-31T00:00:00Z",
		},
		// Case 007
		{
			tic: "2024-01-15T00:00:00Z",
			qnt: +1000,
			res: "2026-10-11T00:00:00Z",
		},
		// Case 008
		{
			tic: "2024-01-15T00:00:00Z",
			qnt: -1000,
			res: "2021-04-20T00:00:00Z",
		},
		// Case 009
		{
			tic: "2024-01-01T00:00:00Z",
			qnt: +365,
			res: "2024-12-31T00:00:00Z",
		},
		// Case 010
		{
			tic: "2024-01-01T00:00:00Z",
			qnt: +366,
			res: "2025-01-01T00:00:00Z",
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
				res = Day(tic.UTC()).Tick(tc.qnt).Time().Format(time.RFC3339Nano)
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}
