package prallel

import (
	"github.com/xh3b4sd/tracer"
	"golang.org/x/sync/errgroup"
)

func Slice[T any](lis []T, fnc func(int, T) error) error {
	var grp errgroup.Group
	{
		grp = errgroup.Group{}
	}

	for i, x := range lis {
		grp.Go(func() error {
			return fnc(i, x)
		})
	}

	{
		err := grp.Wait()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
