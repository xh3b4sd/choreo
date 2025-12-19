package ticker

import "time"

type Interface interface {
	Tick(int) time.Time
}
