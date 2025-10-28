package stream

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/xh3b4sd/choreo/backoff"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Bac is the token backoff used to delay reconnection attempts, which also
	// execute the provided gap callback.
	Bac *backoff.Token

	// Dat is the buffered data channel to which any received data is forwarded.
	// If this channel is not properly buffered or drained, then this blocking
	// channel causes the respectively received websocket messages to be dropped.
	Dat chan<- []byte

	// Dia is the dial factory that returns a generic interface to the underlying
	// websocket connection.
	Dia func() (Connection, error)

	// Gap is the gap callback, which is guaranteed to be executed when the
	// underlying websocket connection got disrupted and re-established. Further,
	// this gap callback is also executed periodically during periods of
	// disconnection, based on the provided token backoff. This function may be
	// used to do work in order to maintain data integrity upon stream
	// interruption.
	//
	//     Sub                                Sub
	//      ↓                                  ↓
	//      [++++++++++++++++++++++] [-------] [+++++++++++]
	//                               ↑   ↑   ↑
	//                              Gap Gap Gap
	//
	Gap func()

	// Sub is the subscription callback, which is guaranteed to be executed when
	// the underlying websocket connection is established. This function may be
	// used to do work in order to configure the established websocket connection.
	Sub func(Connection)
}

type Stream struct {
	bac *backoff.Token
	clo atomic.Bool
	con Connection
	dat chan<- []byte
	dia func() (Connection, error)
	gap func()
	mut sync.Mutex
	sub func(Connection)
}

func New(c Config) *Stream {
	if c.Bac == nil {
		tracer.Panic(fmt.Errorf("%T.Bac must not be empty", c))
	}
	if c.Dat == nil {
		tracer.Panic(fmt.Errorf("%T.Dat must not be empty", c))
	}
	if c.Dia == nil {
		tracer.Panic(fmt.Errorf("%T.Dia must not be empty", c))
	}
	if c.Gap == nil {
		tracer.Panic(fmt.Errorf("%T.Gap must not be empty", c))
	}
	if c.Sub == nil {
		tracer.Panic(fmt.Errorf("%T.Sub must not be empty", c))
	}

	return &Stream{
		bac: c.Bac,
		clo: atomic.Bool{},
		con: nil,
		dat: c.Dat,
		dia: c.Dia,
		gap: c.Gap,
		mut: sync.Mutex{},
		sub: c.Sub,
	}
}
