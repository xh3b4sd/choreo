package success

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/xh3b4sd/tracer"
	"golang.org/x/sync/errgroup"
)

func Test_Success_default(t *testing.T) {
	testCases := []struct {
		suc uint
		des int
	}{
		// Case 000
		{
			suc: 0,
			des: 1,
		},
		// Case 001
		{
			suc: 1,
			des: 1,
		},
		// Case 002
		{
			suc: 2,
			des: 2,
		},
		// Case 003
		{
			suc: 8,
			des: 8,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var mut *Mutex
			{
				mut = New(Config{
					Suc: tc.suc,
				})
			}

			if mut.des != tc.des {
				t.Fatal("expected", tc.des, "got", mut.des)
			}
		})
	}
}

func Test_Success_error(t *testing.T) {
	var mut *Mutex
	{
		mut = New(Config{
			Suc: 2,
		})
	}

	var fai atomic.Int32
	var grp errgroup.Group
	{
		fai = atomic.Int32{}
		grp = errgroup.Group{}
	}

	for range 100 {
		grp.Go(func() error {
			return mut.Success(func() error {
				fai.Add(1)
				return tracer.Mask(testError)
			})
		})
	}

	err := grp.Wait()
	if !isTest(err) {
		t.Fatal("expected", true, "got", err)
	}

	if fai.Load() != 100 {
		t.Fatal("expected", 100, "got", fai.Load())
	}
}

func Test_Success_mutex(t *testing.T) {
	testCases := []struct {
		suc uint
		des int32
	}{
		// Case 000
		{
			suc: 0,
			des: 1,
		},
		// Case 001
		{
			suc: 1,
			des: 1,
		},
		// Case 002
		{
			suc: 2,
			des: 2,
		},
		// Case 003
		{
			suc: 8,
			des: 8,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var mut *Mutex
			{
				mut = New(Config{
					Suc: tc.suc,
				})
			}

			var grp errgroup.Group
			var suc atomic.Int32
			{
				grp = errgroup.Group{}
				suc = atomic.Int32{}
			}

			fnc := func() error {
				suc.Add(1)
				return nil
			}

			for range 100 {
				grp.Go(func() error {
					return mut.Success(fnc)
				})
			}

			err := grp.Wait()
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}

			if suc.Load() != tc.des {
				t.Fatal("expected", tc.des, "got", suc.Load())
			}
		})
	}
}
