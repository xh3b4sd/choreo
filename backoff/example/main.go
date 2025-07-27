package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/xh3b4sd/choreo/backoff"
)

func main() {
	var tok *backoff.Token
	{
		tok = backoff.New(backoff.Config{
			Bac: []time.Duration{
				1 * time.Second,
				2 * time.Second,
				3 * time.Second,
			},
		})
	}

	fnc := func() error {
		fmt.Printf("tried at second %d\n", time.Now().Second())
		return errors.New("error")
	}

	err := tok.Backoff(fnc)
	if err != nil {
		fmt.Printf("%s at second %d\n", err.Error(), time.Now().Second())
	}
}
