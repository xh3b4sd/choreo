package parallel

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var testError = &tracer.Error{
	Description: "This error is only used for testing purposes.",
}

func isTest(err error) bool {
	return errors.Is(err, testError)
}
