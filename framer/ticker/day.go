package ticker

import "time"

type day struct {
	tic time.Time
}

func Day(tic time.Time) Interface {
	return &day{
		tic: tic,
	}
}

func (d *day) Tick(qnt int) time.Time {
	return truDay(d.tic).AddDate(0, 0, qnt)
}

func truDay(tic time.Time) time.Time {
	// Get the base parameters of the current time. Here the things that matter
	// are year, month and day.

	var y, m, d = tic.Date()

	// Set the returned time to the first hour of the truncated day.

	return time.Date(y, m, d, 0, 0, 0, 0, tic.Location())
}
