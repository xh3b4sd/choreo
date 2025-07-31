package success

import "github.com/xh3b4sd/tracer"

func (m *Mutex) Success(fnc func() error) error {
	// Snynchronize calls to Success() so that we control the execution path for
	// exclusive access only. Using a sync.Mutex guarantees only a single
	// execution at once.

	{
		m.mut.Lock()
	}

	// If we executed already successfully to the desired extend, then unlock and
	// return early. The only performance cost for us after all successful
	// executions is to lock and unlock the underlying mutex. For most problem
	// domains, that should not be an issue.

	if m.cur >= m.des {
		m.mut.Unlock()
		return nil
	}

	// At this point we want to execute the injected callback and handle its
	// errors, if any. This code path is only executed the desired amount of
	// times. Note that we do not accumulate the extra costs for defer()
	// statements.

	{
		err := fnc()
		if err != nil {
			m.mut.Unlock()
			return tracer.Mask(err)
		}
	}

	// Increment the successful execution count, so that consecutive calls of
	// Success() can reconcile their own exclusive execution.

	{
		m.cur++
		m.mut.Unlock()
	}

	return nil
}
