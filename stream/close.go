package stream

import (
	"github.com/xh3b4sd/tracer"
)

// Close stops the processing of all websocket messages by disconnecting the
// underlying websocket connection.
func (s *Stream) Close() error {
	{
		s.mut.Lock()
		defer s.mut.Unlock()
	}

	// Ensure that any call to Stream.Close has no effect if we have no valid
	// websocket connection setup yet.

	if s.con == nil {
		return nil
	}

	// Communicate our close state internally so that our read loop failure does
	// not cause reconnection attempts anymore.

	if s.clo.Swap(true) {
		return nil
	}

	err := s.con.Close()
	if err != nil {
		return tracer.Mask(err)
	}

	{
		s.con = nil
	}

	return nil
}
