package parallel

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/xh3b4sd/tracer"
)

func Test_Parallel_Slice_failure(t *testing.T) {
	var inp []int
	var out error
	var num atomic.Int64
	{
		inp = []int{10, 20, 30, 40, 50}
		out = errors.New("test error")
		num = atomic.Int64{}
	}

	fnc := func(_ int, v int) error {
		{
			num.Add(1)
		}

		if v == 30 {
			return tracer.Mask(out)
		}

		return nil
	}

	err := Slice(inp, fnc)
	if !errors.Is(err, out) {
		t.Fatal("expected", out, "got", err)
	}

	if int(num.Load()) != len(inp) {
		t.Fatal("expected", len(inp), "got", num.Load())
	}
}

func Test_Parallel_Slice_success(t *testing.T) {
	var mut sync.Mutex

	var inp []string
	var out []string
	{
		inp = []string{"10", "20", "30", "40", "50"}
		out = make([]string, len(inp))
	}

	fnc := func(i int, x string) error {
		mut.Lock()
		out[i] = x // ensure reliable index/value mapping
		mut.Unlock()
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
