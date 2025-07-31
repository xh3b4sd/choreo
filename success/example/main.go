package main

import (
	"fmt"

	"github.com/xh3b4sd/choreo/success"
	"golang.org/x/sync/errgroup"
)

func main() {
	var mut *success.Mutex
	{
		mut = success.New(success.Config{
			Suc: 2,
		})
	}

	var grp errgroup.Group
	{
		grp = errgroup.Group{}
	}

	for range 100 {
		grp.Go(func() error {
			return mut.Success(func() error {
				fmt.Printf("hello world\n")
				return nil
			})
		})
	}

	err := grp.Wait()
	if err != nil {
		panic(err)
	}
}
