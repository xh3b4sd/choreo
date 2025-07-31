package parallel

import (
	"github.com/xh3b4sd/tracer"
	"golang.org/x/sync/errgroup"
)

// Func is an errgroup.Group wrapper for convenience. Every callback will be
// executed concurrently by the underlying error group. All callbacks execute
// every time even if the very first execution produces an error. The first
// error caught will be returned.
func Func(fnc ...func() error) error {
	var grp errgroup.Group
	{
		grp = errgroup.Group{}
	}

	for _, x := range fnc {
		grp.Go(x)
	}

	{
		err := grp.Wait()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
