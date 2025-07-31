package success

import (
	"sync"
)

type Config struct {
	Suc uint
}

type Mutex struct {
	cur int
	des int
	mut sync.Mutex
}

func New(c Config) *Mutex {
	if c.Suc == 0 {
		c.Suc = 1
	}

	return &Mutex{
		cur: 0,
		des: int(c.Suc),
		mut: sync.Mutex{},
	}
}
