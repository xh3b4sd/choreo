package factory

import (
	"context"
	"time"

	"github.com/coder/websocket"
	"github.com/xh3b4sd/choreo/stream"
	"github.com/xh3b4sd/choreo/stream/connection"
	"github.com/xh3b4sd/tracer"
)

func Coder(add string) func() (stream.Connection, error) {
	return func() (stream.Connection, error) {
		var err error

		var ctx context.Context
		var can context.CancelFunc
		{
			ctx, can = context.WithTimeout(context.Background(), time.Minute)
		}

		{
			defer can()
		}

		var opt *websocket.DialOptions
		{
			opt = &websocket.DialOptions{
				OnPingReceived: func(context.Context, []byte) bool { return true },
			}
		}

		var con *websocket.Conn
		{
			con, _, err = websocket.Dial(ctx, add, opt)
			if err != nil {
				return nil, tracer.Mask(err)
			}
		}

		return &connection.Coder{Con: con}, nil
	}
}
