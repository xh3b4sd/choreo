package ticker

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Framer_Ticker_Month_Tick(t *testing.T) {
	var testCases = []struct {
		now string
		qnt int
		res string
	}{
		// Case 000, Jan -> Jan start
		{
			now: "2024-01-15T12:34:56Z",
			qnt: 0,
			res: "2024-01-01T00:00:00Z",
		},
		// Case 001, Feb (leap year) -> Feb start
		{
			now: "2024-02-29T23:59:59Z",
			qnt: 0,
			res: "2024-02-01T00:00:00Z",
		},
		// Case 002, end-of-month with nanos still snaps to month midnight
		{
			now: "2024-03-31T23:59:59.999999999Z",
			qnt: 0,
			res: "2024-03-01T00:00:00Z",
		},
		// Case 003, Apr -> Apr start
		{
			now: "2024-04-15T12:34:56Z",
			qnt: 0,
			res: "2024-04-01T00:00:00Z",
		},
		// Case 004, boundary stays boundary when qnt=0
		{
			now: "2024-04-01T00:00:00Z",
			qnt: 0,
			res: "2024-04-01T00:00:00Z",
		},
		// Case 005, boundary date with time-of-day still snaps to month midnight
		{
			now: "2024-04-01T12:00:00Z",
			qnt: 0,
			res: "2024-04-01T00:00:00Z",
		},
		// Case 006, boundary date with nanoseconds still snaps to month midnight
		{
			now: "2024-10-01T00:00:00.000000001Z",
			qnt: 0,
			res: "2024-10-01T00:00:00Z",
		},
		// Case 007, +1 month from mid-month
		{
			now: "2024-05-20T08:00:00Z",
			qnt: +1,
			res: "2024-06-01T00:00:00Z",
		},
		// Case 008, -1 month from mid-month
		{
			now: "2024-05-20T08:00:00Z",
			qnt: -1,
			res: "2024-04-01T00:00:00Z",
		},
		// Case 009, +2 months crosses year end (Nov +2 => Jan next year)
		{
			now: "2024-11-20T08:00:00Z",
			qnt: +2,
			res: "2025-01-01T00:00:00Z",
		},
		// Case 010, +12 months same month next year
		{
			now: "2024-05-20T08:00:00Z",
			qnt: +12,
			res: "2025-05-01T00:00:00Z",
		},
		// Case 011, -12 months same month previous year
		{
			now: "2024-05-20T08:00:00Z",
			qnt: -12,
			res: "2023-05-01T00:00:00Z",
		},
		// Case 012, end-of-year snaps to Dec start
		{
			now: "2024-12-31T23:59:59Z",
			qnt: 0,
			res: "2024-12-01T00:00:00Z",
		},
		// Case 013, large positive qnt
		{
			now: "2024-05-20T08:00:00Z",
			qnt: +120,
			res: "2034-05-01T00:00:00Z",
		},
		// Case 014, large negative qnt
		{
			now: "2024-05-20T08:00:00Z",
			qnt: -120,
			res: "2014-05-01T00:00:00Z",
		},
		// Case 015, offset input snaps to month start in same offset
		{
			now: "2024-03-31T23:59:59+02:00",
			qnt: 0,
			res: "2024-03-01T00:00:00+01:00",
		},
		// Case 016, offset input +1 month
		{
			now: "2024-05-20T08:00:00+02:00",
			qnt: +1,
			res: "2024-06-01T00:00:00+02:00",
		},
		// Case 017, offset input -1 month
		{
			now: "2024-05-20T08:00:00+02:00",
			qnt: -1,
			res: "2024-04-01T00:00:00+02:00",
		},
		// Case 018, offset input crosses year end (Nov +2 => Jan next year)
		{
			now: "2024-11-20T08:00:00+01:00",
			qnt: +2,
			res: "2025-01-01T00:00:00+01:00",
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
				res = Month(now).Tick(tc.qnt).Format(time.RFC3339Nano)
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}
