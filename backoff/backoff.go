package backoff

import "github.com/xh3b4sd/tracer"

func (t *Token) Backoff(fnc func() error) error {
	for {
		err := fnc()
		if err != nil {
			// Invoke the sleep function using the currently activated backoff
			// duration.

			{
				t.slp(t.bac[t.ind])
			}

			// Increment the token index, if possible. In case the infinity flag has
			// not been set, and if we are already at the end of the possible token
			// index, return the last error received to terminate the backoff.

			if !t.incTok() && !t.inf {
				return tracer.Mask(err)
			}
		} else {
			// Decrement the token index, if possible. This refills the token index
			// and allows the first backoff durations to be used.

			{
				t.decTok()
			}

			{
				return nil
			}
		}
	}
}

func (t *Token) incTok() bool {
	if t.ind != len(t.bac)-1 {
		t.ind++
		return true
	}

	return false
}

func (t *Token) decTok() bool {
	if t.ind != 0 {
		t.ind--
		return true
	}

	return false
}
