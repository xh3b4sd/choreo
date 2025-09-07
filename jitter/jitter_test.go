package jitter

import (
	"fmt"
	"testing"
	"time"
)

func Test_Jitter_Percent_duration(t *testing.T) {
	testCases := []struct {
		num time.Duration
		per float64
		ran [2]time.Duration
	}{
		// Case 000
		{
			per: 0.1,
			num: 1000000,
			ran: [2]time.Duration{900000, 1100000},
		},
		// Case 001
		{
			per: 0.2,
			num: time.Hour,
			ran: [2]time.Duration{40 * time.Minute, 80 * time.Minute},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var jit *Jitter[time.Duration]
			{
				jit = New[time.Duration](Config{
					Per: tc.per,
				})
			}

			for range 100 {
				var num time.Duration
				{
					num = jit.Percent(tc.num)
				}

				if num == tc.num {
					t.Fatal("expected !=", tc.num, "got", num)
				}
				if num < tc.ran[0] {
					t.Fatal("expected <", tc.ran[0], "got", num)
				}
				if num > tc.ran[1] {
					t.Fatal("expected >", tc.ran[1], "got", num)
				}
			}
		})
	}
}

func Test_Jitter_Percent_float64(t *testing.T) {
	testCases := []struct {
		num float64
		per float64
		ran [2]float64
	}{
		// Case 000
		{
			per: 0.1,
			num: 100,
			ran: [2]float64{90, 110},
		},
		// Case 001
		{
			per: 0.2,
			num: 60,
			ran: [2]float64{40, 80},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var jit *Jitter[float64]
			{
				jit = New[float64](Config{
					Per: tc.per,
				})
			}

			for range 100 {
				var num float64
				{
					num = jit.Percent(tc.num)
				}

				if num == tc.num {
					t.Fatal("expected !=", tc.num, "got", num)
				}
				if num < tc.ran[0] {
					t.Fatal("expected <", tc.ran[0], "got", num)
				}
				if num > tc.ran[1] {
					t.Fatal("expected >", tc.ran[1], "got", num)
				}
			}
		})
	}
}
