package parallel

import (
	"sync/atomic"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/xh3b4sd/tracer"
)

func Test_Parallel_Slice_failure(t *testing.T) {
	var inp []int
	var num atomic.Int64
	{
		inp = []int{10, 20, 30, 40, 50}
		num = atomic.Int64{}
	}

	fnc := func(_ int, v int) error {
		{
			num.Add(1)
		}

		if v == 30 {
			return tracer.Mask(testError)
		}

		return nil
	}

	err := Slice(inp, fnc)
	if !isTest(err) {
		t.Fatal("expected", true, "got", err)
	}

	if int(num.Load()) != len(inp) {
		t.Fatal("expected", len(inp), "got", num.Load())
	}
}

// Test_Parallel_Slice_success ensures reliable index/value mapping so that
// distinct items of the same slice can be manipulated concurrently without
// further synchronization.
func Test_Parallel_Slice_success(t *testing.T) {
	var inp []string
	var out []string
	{
		inp = []string{"10", "20", "30", "40", "50"}
		out = make([]string, len(inp))
	}

	fnc := func(i int, x string) error {
		out[i] = x
		return nil
	}

	err := Slice(inp, fnc)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if dif := cmp.Diff(inp, out); dif != "" {
		t.Fatalf("-expected +actual:\n%s", dif)
	}
}
