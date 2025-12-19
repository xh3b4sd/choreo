package ticker

import "time"

type quarter struct {
	tic time.Time
}

func Quarter(tic time.Time) Interface {
	return &quarter{
		tic: tic,
	}
}

func (q *quarter) Tick(qnt int) time.Time {
	return truQua(q.tic).AddDate(0, 3*qnt, 0)
}

func truQua(tic time.Time) time.Time {
	// Get the base parameters of the current time. Here the things that matter
	// are year and month.

	var y, m, _ = tic.Date()

	// Truncate and scale up the month index by factors of 3, because 3 months
	// equal 1 quarter.

	var q = time.Month(((int(m)-1)/3)*3 + 1)

	// Set the returned time to the first day of the truncated quarter.

	return time.Date(y, q, 1, 0, 0, 0, 0, tic.Location())
}
