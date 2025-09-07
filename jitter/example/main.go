package main

import (
	"fmt"
	"time"

	"github.com/xh3b4sd/choreo/jitter"
)

func main() {
	jit := jitter.New[time.Duration](jitter.Config{
		Per: 0.1,
	})

	for range 5 {
		fmt.Printf("%s\n", jit.Percent(time.Minute))
	}
}
