package sequence

// Wrap is a wrapper for Func so that Func is not immediatelly executed, but
// instead it is left to the caller to execute the returned wrapper function.
func Wrap(fnc ...func() error) func() error {
	return func() error {
		return Func(fnc...)
	}
}
