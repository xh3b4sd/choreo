package parallel

import (
	"errors"
	"slices"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/xh3b4sd/tracer"
)

func Test_Parallel_Slice_failure(t *testing.T) {
	var inp []int
	var out error
	{
		inp = []int{10, 20, 30, 40, 50}
		out = errors.New("test error")
	}

	fnc := func(_ int, v int) error {
		if v == 30 {
			return tracer.Mask(out)
		}
		return nil
	}

	err := Slice(inp, fnc)
	if !errors.Is(err, out) {
		t.Fatal("expected", out, "got", err)
	}
}

func Test_Parallel_Slice_success(t *testing.T) {
	var mut sync.Mutex

	var inp []string
	var out []string
	{
		inp = []string{"10", "20", "30", "40", "50"}
	}

	fnc := func(_ int, v string) error {
		mut.Lock()
		out = append(out, v)
		mut.Unlock()
		return nil
	}

	err := Slice(inp, fnc)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	{
		slices.Sort(inp)
		slices.Sort(out)
	}

	if dif := cmp.Diff(inp, out); dif != "" {
		t.Fatalf("-expected +actual:\n%s", dif)
	}
}
