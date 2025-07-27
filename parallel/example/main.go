package main

import (
	"fmt"

	"github.com/xh3b4sd/choreo/parallel"
)

func main() {
	var inp []string
	{
		inp = []string{"10", "20", "30", "40", "50"}
	}

	fnc := func(i int, x string) error {
		fmt.Printf("index %d value %s\n", i, x)
		return nil
	}

	err := parallel.Slice(inp, fnc)
	if err != nil {
		panic(err)
	}
}
