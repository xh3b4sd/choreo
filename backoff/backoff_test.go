package backoff

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Backoff_retry(t *testing.T) {
	testCases := []struct {
		des int
		inf bool
		act int
		ind []int
		mat func(err error) bool
	}{
		// Case 000, no retries, 1 call, no token index
		{
			des: 1,
			inf: false,
			act: 1,
			ind: nil,
			mat: isNil,
		},
		// Case 001, 1 retry, 2 calls, token index 0
		{
			des: 2,
			inf: false,
			act: 2,
			ind: []int{0},
			mat: isNil,
		},
		// Case 002, 2 retries, 3 calls, token index 0 1
		{
			des: 3,
			inf: false,
			act: 3,
			ind: []int{0, 1},
			mat: isNil,
		},
		// Case 003, 3 retries, 4 calls, token index 0 1 2
		{
			des: 4,
			inf: false,
			act: 4,
			ind: []int{0, 1, 2},
			mat: isNil,
		},
		// Case 004, 4 retries, 5 calls, token index 0 1 2 3
		{
			des: 5,
			inf: false,
			act: 4,
			ind: []int{0, 1, 2, 3},
			mat: isTest,
		},
		// Case 005, 4 retries, 5 calls, token index 0 1 2 3
		{
			des: 6,
			inf: false,
			act: 4,
			ind: []int{0, 1, 2, 3},
			mat: isTest,
		},
		// Case 006, 6 retries, 7 calls, token index 0 1 2 3 3 3 where the last 2
		// backoff durations are activated based on the infinity flag
		{
			des: 7,
			inf: true,
			act: 7,
			ind: []int{0, 1, 2, 3, 3, 3},
			mat: isNil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var tok *Token
			{
				tok = New(Config{
					Bac: []time.Duration{0, 1, 2, 3},
					Inf: tc.inf,
				})
			}

			var ind []int
			tok.slp = func(dur time.Duration) {
				ind = append(ind, int(dur))
			}

			var act int
			fnc := func(des int) func() error {
				return func() error {
					{
						act++
					}

					if act < des {
						return testError
					}

					return nil
				}
			}

			err := tok.Backoff(fnc(tc.des))
			if !tc.mat(err) {
				t.Fatal("expected", true, "got", err)
			}

			if act != tc.act {
				t.Fatal("expected", tc.act, "got", act)
			}

			if dif := cmp.Diff(tc.ind, ind); dif != "" {
				t.Fatalf("-expected +actual:\n%s", dif)
			}
		})
	}
}

// Test_Backoff_replenish verifies that successful executions refill the
// exhausted error budget, by moving the token index back again towards zero,
// but never past zero itself. In other words, the beginning of the configured
// backoff durations will be activated again at token index zero.
func Test_Backoff_refill(t *testing.T) {
	var err error

	suc := func() func() error {
		return func() error {
			return nil
		}
	}

	fai := func() func() error {
		var cal int

		return func() error {
			{
				cal++
			}

			if cal <= 2 {
				return testError
			}

			return nil
		}
	}

	var tok *Token
	{
		tok = New(Config{
			Bac: []time.Duration{0, 1, 2, 3},
		})
	}

	{
		if tok.ind != 0 {
			t.Fatal("expected", 0, "got", tok.ind)
		}
	}

	{
		err = tok.Backoff(fai())
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}

	{
		if tok.ind != 1 {
			t.Fatal("expected", 1, "got", tok.ind)
		}
	}

	{
		err = tok.Backoff(fai())
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}

	{
		if tok.ind != 2 {
			t.Fatal("expected", 2, "got", tok.ind)
		}
	}

	{
		err = tok.Backoff(suc())
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}

	{
		if tok.ind != 1 {
			t.Fatal("expected", 1, "got", tok.ind)
		}
	}

	{
		err = tok.Backoff(fai())
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}

	{
		if tok.ind != 2 {
			t.Fatal("expected", 2, "got", tok.ind)
		}
	}

	{
		err = tok.Backoff(suc())
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}

	{
		if tok.ind != 1 {
			t.Fatal("expected", 1, "got", tok.ind)
		}
	}

	{
		err = tok.Backoff(suc())
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}

	{
		if tok.ind != 0 {
			t.Fatal("expected", 0, "got", tok.ind)
		}
	}

	// Multiple successful executions do not move the token index past zero into
	// the negative. This ensures that the very first token indexed backoff
	// duration will be respected on the next occuring execution failure.

	{
		err = tok.Backoff(suc())
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}

	{
		if tok.ind != 0 {
			t.Fatal("expected", 0, "got", tok.ind)
		}
	}

	{
		err = tok.Backoff(suc())
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}

	{
		if tok.ind != 0 {
			t.Fatal("expected", 0, "got", tok.ind)
		}
	}

	// Starting off again from the token index base, the next execution failure
	// causes the second token index to be activated again.

	{
		err = tok.Backoff(fai())
		if err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}

	{
		if tok.ind != 1 {
			t.Fatal("expected", 1, "got", tok.ind)
		}
	}
}
