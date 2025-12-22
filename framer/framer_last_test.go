package framer

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Framer_Last_Day(t *testing.T) {
	var testCases = []struct {
		min string
		max string
		qnt int
		res []string
	}{
		// Case 000, qnt=1 crosses year boundary
		{
			min: "2024-12-30T00:00:00Z",
			max: "2025-01-03T00:00:00Z",
			qnt: 1,
			res: []string{
				"2024-12-30T00:00:00Z",
				"2024-12-31T00:00:00Z",
				"2025-01-01T00:00:00Z",
				"2025-01-02T00:00:00Z",
			},
		},
		// Case 001, qnt=2 crosses year boundary
		{
			min: "2024-12-30T00:00:00Z",
			max: "2025-01-05T00:00:00Z",
			qnt: 2,
			res: []string{
				"2024-12-30T00:00:00Z",
				"2025-01-01T00:00:00Z",
				"2025-01-03T00:00:00Z",
			},
		},
		// Case 002, qnt=1 includes leap day
		{
			min: "2024-02-28T00:00:00Z",
			max: "2024-03-02T00:00:00Z",
			qnt: 1,
			res: []string{
				"2024-02-28T00:00:00Z",
				"2024-02-29T00:00:00Z",
				"2024-03-01T00:00:00Z",
			},
		},
		// Case 003, spans multiple years with larger qnt
		{
			min: "2023-12-25T00:00:00Z",
			max: "2024-01-10T00:00:00Z",
			qnt: 7,
			res: []string{
				"2023-12-25T00:00:00Z",
				"2024-01-01T00:00:00Z",
				"2024-01-08T00:00:00Z",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var min time.Time
			{
				min, err = time.Parse(time.RFC3339Nano, tc.min)
				if err != nil {
					t.Fatal(err)
				}
			}

			var max time.Time
			{
				max, err = time.Parse(time.RFC3339Nano, tc.max)
				if err != nil {
					t.Fatal(err)
				}
			}

			var fra *Framer
			{
				fra = New(Config{
					Min: min,
					Max: max,
				})
			}

			var res []string

			for t := fra.Tick(); !fra.Last(); fra.Day(tc.qnt) {
				res = append(res, t.Time().Format(time.RFC3339Nano))
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}

func Test_Framer_Last_Month(t *testing.T) {
	var testCases = []struct {
		min string
		max string
		qnt int
		res []string
	}{
		// Case 000, qnt=1 crosses year boundary
		{
			min: "2024-10-01T00:00:00Z",
			max: "2025-03-01T00:00:00Z",
			qnt: 1,
			res: []string{
				"2024-10-01T00:00:00Z",
				"2024-11-01T00:00:00Z",
				"2024-12-01T00:00:00Z",
				"2025-01-01T00:00:00Z",
				"2025-02-01T00:00:00Z",
			},
		},
		// Case 001, qnt=2 crosses year boundary
		{
			min: "2024-10-01T00:00:00Z",
			max: "2025-04-01T00:00:00Z",
			qnt: 2,
			res: []string{
				"2024-10-01T00:00:00Z",
				"2024-12-01T00:00:00Z",
				"2025-02-01T00:00:00Z",
			},
		},
		// Case 002, spans multiple years
		{
			min: "2023-11-01T00:00:00Z",
			max: "2025-04-01T00:00:00Z",
			qnt: 5,
			res: []string{
				"2023-11-01T00:00:00Z",
				"2024-04-01T00:00:00Z",
				"2024-09-01T00:00:00Z",
				"2025-02-01T00:00:00Z",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var min time.Time
			{
				min, err = time.Parse(time.RFC3339Nano, tc.min)
				if err != nil {
					t.Fatal(err)
				}
			}

			var max time.Time
			{
				max, err = time.Parse(time.RFC3339Nano, tc.max)
				if err != nil {
					t.Fatal(err)
				}
			}

			var fra *Framer
			{
				fra = New(Config{
					Min: min,
					Max: max,
				})
			}

			var res []string

			for t := fra.Tick(); !fra.Last(); fra.Month(tc.qnt) {
				res = append(res, t.Time().Format(time.RFC3339Nano))
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}

func Test_Framer_Last_Quarter(t *testing.T) {
	var testCases = []struct {
		min string
		max string
		qnt int
		res []string
	}{
		// Case 000, qnt=1 within a single year
		{
			min: "2024-01-01T00:00:00Z",
			max: "2025-01-01T00:00:00Z",
			qnt: 3,
			res: []string{
				"2024-01-01T00:00:00Z",
				"2024-04-01T00:00:00Z",
				"2024-07-01T00:00:00Z",
				"2024-10-01T00:00:00Z",
			},
		},
		// Case 001, qnt=2 within a single year
		{
			min: "2024-01-01T00:00:00Z",
			max: "2025-01-01T00:00:00Z",
			qnt: 6,
			res: []string{
				"2024-01-01T00:00:00Z",
				"2024-07-01T00:00:00Z",
			},
		},
		// Case 002, overlaps multiple years
		{
			min: "2023-10-01T00:00:00Z",
			max: "2025-07-01T00:00:00Z",
			qnt: 3,
			res: []string{
				"2023-10-01T00:00:00Z",
				"2024-01-01T00:00:00Z",
				"2024-04-01T00:00:00Z",
				"2024-07-01T00:00:00Z",
				"2024-10-01T00:00:00Z",
				"2025-01-01T00:00:00Z",
				"2025-04-01T00:00:00Z",
			},
		},
		// Case 003, overlaps multiple years with qnt=3
		{
			min: "2023-10-01T00:00:00Z",
			max: "2025-10-01T00:00:00Z",
			qnt: 9,
			res: []string{
				"2023-10-01T00:00:00Z",
				"2024-07-01T00:00:00Z",
				"2025-04-01T00:00:00Z",
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var min time.Time
			{
				min, err = time.Parse(time.RFC3339Nano, tc.min)
				if err != nil {
					t.Fatal(err)
				}
			}

			var max time.Time
			{
				max, err = time.Parse(time.RFC3339Nano, tc.max)
				if err != nil {
					t.Fatal(err)
				}
			}

			var fra *Framer
			{
				fra = New(Config{
					Min: min,
					Max: max,
				})
			}

			var res []string

			for t := fra.Tick(); !fra.Last(); fra.Month(tc.qnt) {
				res = append(res, t.Time().Format(time.RFC3339Nano))
			}

			if dif := cmp.Diff(tc.res, res); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}
