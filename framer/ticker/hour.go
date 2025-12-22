package ticker

import "time"

type hour struct {
	tic time.Time
}

func Hour(tic time.Time) Interface {
	return &hour{
		tic: tic,
	}
}

func (h *hour) Round(mul int) Interface {
	{
		h.tic = rndHou(h.tic, mul)
	}

	return h
}

func (h *hour) Tick(qnt int) Interface {
	{
		h.tic = h.tic.Add(time.Duration(qnt) * time.Hour)
	}

	return h
}

func (h *hour) Time() time.Time {
	return h.tic
}

func rndHou(tic time.Time, mul int) time.Time {
	var yea, mon, day = tic.Date()

	var zer = time.Date(yea, mon, day, 0, 0, 0, 0, tic.Location())

	return zer.Add(tic.Sub(zer).Round(time.Duration(mul) * time.Hour))
}
