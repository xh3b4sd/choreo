package sequence

import (
	"github.com/xh3b4sd/tracer"
)

// Func is a loop wrapper for convenience. Every callback will be executed
// sequentially by the underlying for loop. All callbacks execute every time
// until the very first execution produces an error. The first error caught will
// be returned.
func Func(fnc ...func() error) error {
	for _, x := range fnc {
		err := x()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
