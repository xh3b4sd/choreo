package parallel

import (
	"sync/atomic"
	"testing"

	"github.com/xh3b4sd/tracer"
)

func Test_Parallel_Func(t *testing.T) {
	var num atomic.Int64
	{
		num = atomic.Int64{}
	}

	fn1 := func() error {
		num.Add(1)
		return nil
	}

	fn2 := func() error {
		num.Add(1)
		return nil
	}

	fn3 := func() error {
		num.Add(1)
		return tracer.Mask(testError)
	}

	fn4 := func() error {
		num.Add(1)
		return nil
	}

	fn5 := func() error {
		num.Add(1)
		return nil
	}

	err := Func(fn1, fn2, fn3, fn4, fn5)
	if !isTest(err) {
		t.Fatal("expected", true, "got", err)
	}

	if num.Load() != 5 {
		t.Fatal("expected", 5, "got", num.Load())
	}
}

func Test_Parallel_Wrap(t *testing.T) {
	var num atomic.Int64
	{
		num = atomic.Int64{}
	}

	fn1 := func() error {
		num.Add(1)
		return nil
	}

	fn2 := func() error {
		num.Add(1)
		return nil
	}

	fn3 := func() error {
		num.Add(1)
		return tracer.Mask(testError)
	}

	fn4 := func() error {
		num.Add(1)
		return nil
	}

	fn5 := func() error {
		num.Add(1)
		return nil
	}

	var fnc func() error
	{
		fnc = Wrap(fn1, fn2, fn3, fn4, fn5)
	}

	err := fnc()
	if !isTest(err) {
		t.Fatal("expected", true, "got", err)
	}

	if num.Load() != 5 {
		t.Fatal("expected", 5, "got", num.Load())
	}
}
