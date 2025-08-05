package sequence

import (
	"testing"

	"github.com/xh3b4sd/tracer"
)

// Test_Sequence_Func verifies that sequence.Func executes sequentially and
// stops at first error. Running this test using the -race flag guarantees that
// there is no concurrency involved, because the map access inside the test
// functions is not synchronized.
func Test_Sequence_Func(t *testing.T) {
	var num map[int]struct{}
	{
		num = make(map[int]struct{})
	}

	fn1 := func() error {
		num[1] = struct{}{}
		return nil
	}

	fn2 := func() error {
		num[2] = struct{}{}
		return nil
	}

	fn3 := func() error {
		num[3] = struct{}{}
		return tracer.Mask(testError)
	}

	fn4 := func() error {
		num[4] = struct{}{}
		return nil
	}

	fn5 := func() error {
		num[5] = struct{}{}
		return nil
	}

	err := Func(fn1, fn2, fn3, fn4, fn5)
	if !isTest(err) {
		t.Fatal("expected", true, "got", err)
	}

	if len(num) != 3 {
		t.Fatal("expected", 4, "got", len(num))
	}
	if _, exi := num[1]; !exi {
		t.Fatal("expected", true, "got", false)
	}
	if _, exi := num[2]; !exi {
		t.Fatal("expected", true, "got", false)
	}
	if _, exi := num[3]; !exi {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Sequence_Wrap(t *testing.T) {
	var num map[int]struct{}
	{
		num = make(map[int]struct{})
	}

	fn1 := func() error {
		num[1] = struct{}{}
		return nil
	}

	fn2 := func() error {
		num[2] = struct{}{}
		return nil
	}

	fn3 := func() error {
		num[3] = struct{}{}
		return tracer.Mask(testError)
	}

	fn4 := func() error {
		num[4] = struct{}{}
		return nil
	}

	fn5 := func() error {
		num[5] = struct{}{}
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

	if len(num) != 3 {
		t.Fatal("expected", 4, "got", len(num))
	}
	if _, exi := num[1]; !exi {
		t.Fatal("expected", true, "got", false)
	}
	if _, exi := num[2]; !exi {
		t.Fatal("expected", true, "got", false)
	}
	if _, exi := num[3]; !exi {
		t.Fatal("expected", true, "got", false)
	}
}
