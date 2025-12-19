package ticker

import "time"

type month struct {
	tic time.Time
}

func Month(tic time.Time) Interface {
	return &month{
		tic: tic,
	}
}

func (m *month) Tick(qnt int) time.Time {
	return truMon(m.tic).AddDate(0, qnt, 0)
}

func truMon(tic time.Time) time.Time {
	// Get the base parameters of the current time. Here the things that matter
	// are year and month.

	var y, m, _ = tic.Date()

	// Set the returned time to the first day of the truncated month.

	return time.Date(y, m, 1, 0, 0, 0, 0, tic.Location())
}
