package main

import (
	"fmt"
	"time"

	"github.com/xh3b4sd/choreo/framer"
)

func main() {
	var fra *framer.Framer
	{
		fra = framer.New(framer.Config{
			Min: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			Max: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
		})
	}

	for t := fra.Tick(); !fra.Last(); fra.Month(+3) {
		var sta time.Time
		var end time.Time
		{
			sta = t.Time()
			end = t.Month(+3)
		}

		{
			fmt.Printf("%s\t%s\n", sta.String(), end.String())
		}
	}
}
