package stream

import (
	"github.com/xh3b4sd/tracer"
)

// Daemon manages the underlying websocket connections in order to forward any
// websocket message to the provided data channel.
func (s *Stream) Daemon() error {
	fnc := func() error {
		// If a websocket connection is set and if Stream.dial is effectively called
		// again, then we want to execute the gap callback, potentially for the last
		// time, if this next connection attempt ends up being successful.

		var gap bool
		{
			s.mut.Lock()
			gap = s.con != nil
			s.mut.Unlock()
		}

		if gap {
			s.gap()
		}

		// Try to establish the websocket connection, either the first time, or just
		// again and again.

		err := s.dial()
		if err != nil {
			return tracer.Mask(err)
		}

		return nil
	}

	err := s.bac.Backoff(fnc)
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}

func (s *Stream) dial() error {
	var err error

	var con Connection
	{
		con, err = s.dia()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		s.mut.Lock()
		s.con = con
		s.mut.Unlock()
	}

	// Allow the caller to setup the established websocket connection by executing
	// the given subscription callback.

	{
		s.sub(con)
	}

	for {
		err = s.read(con)
		if err != nil {
			return tracer.Mask(err)
		}
	}
}

func (s *Stream) read(con Connection) error {
	dat, err := con.Read()
	if err != nil {
		// If the error that we just received presumably originates from our
		// internal closing mechanism, then return nil in order to stop all
		// further processing.

		if s.clo.Load() {
			return nil
		}

		return tracer.Mask(err)
	}

	select {
	case s.dat <- dat:
		// forward data
	default:
		// drop message
	}

	return nil
}
