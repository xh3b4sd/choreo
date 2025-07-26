package parallel

import (
	"github.com/xh3b4sd/tracer"
	"golang.org/x/sync/errgroup"
)

// Slice is a generic errgroup.Group wrapper for data lists. Every item of the
// given list is passed to the provided callback, and all resulting callbacks
// are being executed concurrently by the underlying error group. All callbacks
// execute every time even if the very first execution produces an error. The
// first error caught will be returned.
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
