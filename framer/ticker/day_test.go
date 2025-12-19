package ticker

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Framer_Ticker_Day_Tick(t *testing.T) {
	var testCases = []struct {
		now string
		qnt int
		res string
	}{
		// Case 000, snaps to day midnight
		{
			now: "2024-01-15T12:34:56Z",
			qnt: 0,
			res: "2024-01-15T00:00:00Z",
		},
		// Case 001, boundary stays boundary when qnt=0
		{
			now: "2024-01-15T00:00:00Z",
			qnt: 0,
			res: "2024-01-15T00:00:00Z",
		},
		// Case 002, boundary with nanoseconds still snaps to midnight
		{
			now: "2024-01-15T00:00:00.000000001Z",
			qnt: 0,
			res: "2024-01-15T00:00:00Z",
		},
		// Case 003, end-of-day with nanos still snaps to midnight
		{
			now: "2024-01-15T23:59:59.999999999Z",
			qnt: 0,
			res: "2024-01-15T00:00:00Z",
		},
		// Case 004, +1 day
		{
			now: "2024-01-15T12:34:56Z",
			qnt: +1,
			res: "2024-01-16T00:00:00Z",
		},
		// Case 005, -1 day
		{
			now: "2024-01-15T12:34:56Z",
			qnt: -1,
			res: "2024-01-14T00:00:00Z",
		},
		// Case 006, leap day snaps to midnight
		{
			now: "2024-02-29T23:59:59Z",
			qnt: 0,
			res: "2024-02-29T00:00:00Z",
		},
		// Case 007, leap day +1 crosses to March 1
		{
			now: "2024-02-29T12:00:00Z",
			qnt: +1,
			res: "2024-03-01T00:00:00Z",
		},
		// Case 008, Dec 31 +1 crosses year
		{
			now: "2024-12-31T23:59:59Z",
			qnt: +1,
			res: "2025-01-01T00:00:00Z",
		},
		// Case 009, Jan 1 -1 crosses to previous year
		{
			now: "2024-01-01T12:00:00Z",
			qnt: -1,
			res: "2023-12-31T00:00:00Z",
		},
		// Case 010, large positive qnt
		{
			now: "2024-01-15T12:34:56Z",
			qnt: +365,
			res: "2025-01-14T00:00:00Z",
		},
		// Case 011, large negative qnt
		{
			now: "2024-01-15T12:34:56Z",
			qnt: -365,
			res: "2023-01-15T00:00:00Z",
		},
		// Case 012, offset input snaps to day midnight in same offset
		{
			now: "2024-03-31T23:59:59+02:00",
			qnt: 0,
			res: "2024-03-31T00:00:00Z",
		},
		// Case 013, offset input +1 day
		{
			now: "2024-03-31T23:59:59+02:00",
			qnt: +1,
			res: "2024-04-01T00:00:00Z",
		},
		// Case 014, offset input -1 day
		{
			now: "2024-03-31T23:59:59+02:00",
			qnt: -1,
			res: "2024-03-30T00:00:00Z",
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
				res = Day(now.UTC()).Tick(tc.qnt).Format(time.RFC3339Nano)
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}
