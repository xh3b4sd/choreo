package ticker

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Framer_Ticker_Quarter_Tick(t *testing.T) {
	var testCases = []struct {
		now string
		qnt int
		res string
	}{
		// Case 000, Jan -> Q1 start
		{
			now: "2024-01-15T12:34:56Z",
			qnt: 0,
			res: "2024-01-01T00:00:00Z",
		},
		// Case 001, qnt=0 returns quarter start
		{
			now: "2024-08-15T12:00:00Z",
			qnt: 0,
			res: "2024-07-01T00:00:00Z",
		},
		// Case 002, Feb -> Q1 start (leap day)
		{
			now: "2024-02-29T23:59:59Z",
			qnt: 0,
			res: "2024-01-01T00:00:00Z",
		},
		// Case 003, Mar -> Q1 start
		{
			now: "2024-03-31T23:59:59.999999999Z",
			qnt: 0,
			res: "2024-01-01T00:00:00Z",
		},
		// Case 004, Apr -> Q2 start
		{
			now: "2024-04-15T12:34:56Z",
			qnt: 0,
			res: "2024-04-01T00:00:00Z",
		},
		// Case 005, May -> Q2 start
		{
			now: "2024-05-15T12:34:56Z",
			qnt: 0,
			res: "2024-04-01T00:00:00Z",
		},
		// Case 006, Jun -> Q2 start
		{
			now: "2024-06-30T23:59:59Z",
			qnt: 0,
			res: "2024-04-01T00:00:00Z",
		},
		// Case 007, Jul -> Q3 start
		{
			now: "2024-07-15T12:34:56Z",
			qnt: 0,
			res: "2024-07-01T00:00:00Z",
		},
		// Case 008, Aug -> Q3 start
		{
			now: "2024-08-15T12:34:56Z",
			qnt: 0,
			res: "2024-07-01T00:00:00Z",
		},
		// Case 009, Sep -> Q3 start
		{
			now: "2024-09-30T23:59:59Z",
			qnt: 0,
			res: "2024-07-01T00:00:00Z",
		},
		// Case 010, Oct -> Q4 start
		{
			now: "2024-10-15T12:34:56Z",
			qnt: 0,
			res: "2024-10-01T00:00:00Z",
		},
		// Case 011, Nov -> Q4 start
		{
			now: "2024-11-15T12:34:56Z",
			qnt: 0,
			res: "2024-10-01T00:00:00Z",
		},
		// Case 012, Dec -> Q4 start
		{
			now: "2024-12-31T23:59:59Z",
			qnt: 0,
			res: "2024-10-01T00:00:00Z",
		},
		// Case 013, boundary stays boundary when qnt=0
		{
			now: "2024-04-01T00:00:00Z",
			qnt: 0,
			res: "2024-04-01T00:00:00Z",
		},
		// Case 014, boundary +1 quarter
		{
			now: "2024-04-01T00:00:00Z",
			qnt: +1,
			res: "2024-07-01T00:00:00Z",
		},
		// Case 015, mid-quarter +1 quarter
		{
			now: "2024-05-20T08:00:00Z",
			qnt: +1,
			res: "2024-07-01T00:00:00Z",
		},
		// Case 016,
		{
			now: "2024-08-15T12:00:00Z",
			qnt: +1,
			res: "2024-10-01T00:00:00Z",
		},
		// Case 017, +2 quarters crosses year
		{
			now: "2024-11-20T08:00:00Z",
			qnt: +2,
			res: "2025-04-01T00:00:00Z",
		},
		// Case 018, large qnt
		{
			now: "2024-05-20T08:00:00Z",
			qnt: +8,
			res: "2026-04-01T00:00:00Z",
		},
		// Case 019, mid Q1 -1 => previous Q4
		{
			now: "2024-02-10T12:00:00Z",
			qnt: -1,
			res: "2023-10-01T00:00:00Z",
		},
		// Case 020, boundary -1
		{
			now: "2024-07-01T00:00:00Z",
			qnt: -1,
			res: "2024-04-01T00:00:00Z",
		},
		// Case 021, end of year -1 => Q3 start
		{
			now: "2024-12-31T23:59:59Z",
			qnt: -1,
			res: "2024-07-01T00:00:00Z",
		},
		// Case 022, negative qnt (acts like subtract)
		{
			now: "2024-05-20T08:00:00Z",
			qnt: -1,
			res: "2024-01-01T00:00:00Z",
		},
		// Case 023, -2 quarters across year boundary
		{
			now: "2024-01-15T12:00:00Z",
			qnt: -2,
			res: "2023-07-01T00:00:00Z",
		},
		// Case 024, boundary date with time-of-day still snaps to quarter midnight
		{
			now: "2024-04-01T12:00:00Z",
			qnt: 0,
			res: "2024-04-01T00:00:00Z",
		},
		// Case 025, boundary date with nanoseconds still snaps to quarter midnight
		{
			now: "2024-10-01T00:00:00.000000001Z",
			qnt: 0,
			res: "2024-10-01T00:00:00Z",
		},
		// Case 026, qnt=0 on exact quarter boundary returns the same boundary
		{
			now: "2024-07-01T00:00:00Z",
			qnt: 0,
			res: "2024-07-01T00:00:00Z",
		},
		// Case 027, qnt=-1 from exact quarter boundary steps back one quarter
		{
			now: "2024-07-01T00:00:00Z",
			qnt: -1,
			res: "2024-04-01T00:00:00Z",
		},
		// Case 028, qnt=1 from exact quarter boundary steps forward one quarter
		{
			now: "2024-07-01T00:00:00Z",
			qnt: +1,
			res: "2024-10-01T00:00:00Z",
		},
		// Case 029, end-of-quarter with offset snaps to quarter start in same offset
		{
			now: "2024-03-31T23:59:59+02:00",
			qnt: 0,
			res: "2024-01-01T00:00:00Z",
		},
		// Case 030, mid-quarter with offset +1 quarter preserves offset boundary
		{
			now: "2024-05-20T08:00:00+02:00",
			qnt: +1,
			res: "2024-07-01T00:00:00Z",
		},
		// Case 031, qnt=-1 with offset steps back one quarter
		{
			now: "2024-05-20T08:00:00+02:00",
			qnt: -1,
			res: "2024-01-01T00:00:00Z",
		},
		// Case 032, very large positive qnt
		{
			now: "2024-05-20T08:00:00Z",
			qnt: +40,
			res: "2034-04-01T00:00:00Z",
		},
		// Case 033, very large negative qnt
		{
			now: "2024-05-20T08:00:00Z",
			qnt: -40,
			res: "2014-04-01T00:00:00Z",
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
				res = Quarter(now.UTC()).Tick(tc.qnt).Format(time.RFC3339Nano)
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}
