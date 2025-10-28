package connection

import (
	"context"
	"time"

	"github.com/coder/websocket"
	"github.com/xh3b4sd/tracer"
)

type Coder struct {
	Con *websocket.Conn
}

func (c *Coder) Close() error {
	err := c.Con.Close(websocket.StatusNormalClosure, "")
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}

func (c *Coder) Read() ([]byte, error) {
	var ctx context.Context
	var can context.CancelFunc
	{
		ctx, can = context.WithTimeout(context.Background(), 15*time.Second)
	}

	{
		defer can()
	}

	_, byt, err := c.Con.Read(ctx)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return byt, nil
}

func (c *Coder) Write(byt []byte) error {
	var ctx context.Context
	var can context.CancelFunc
	{
		ctx, can = context.WithTimeout(context.Background(), 15*time.Second)
	}

	{
		defer can()
	}

	err := c.Con.Write(ctx, websocket.MessageText, byt)
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}
