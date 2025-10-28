package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/xh3b4sd/choreo/backoff"
	"github.com/xh3b4sd/choreo/stream"
	"github.com/xh3b4sd/choreo/stream/factory"
	"github.com/xh3b4sd/tracer"
)

func main() {
	var add string
	{
		add = fmt.Sprintf("wss://stream.coingecko.com/v1?%s=%s", "x_cg_pro_api_key", os.Getenv("COINGECKO_API_KEY"))
	}

	var bac *backoff.Token
	{
		bac = backoff.New(backoff.Config{
			Bac: []time.Duration{
				1 * time.Second,
				2 * time.Second,
				3 * time.Second,
			},
			Inf: true,
		})
	}

	var dat chan []byte
	{
		dat = make(chan []byte, 100)
	}

	var str *stream.Stream
	{
		str = stream.New(stream.Config{
			Bac: bac,
			Dat: dat,
			Dia: factory.Coder(add),
			Gap: func() {
				fmt.Printf("GAP\n")
			},
			Sub: func(c stream.Connection) {
				fmt.Printf("SUB\n")

				// 1. Subscribe to the CGSimplePrice channel.
				sub := []byte(`{
          "command": "subscribe",
          "identifier": "{\"channel\":\"CGSimplePrice\"}"
        }`)

				err := c.Write(sub)
				if err != nil {
					tracer.Panic(err)
				}

				// 2. Tell the server which coin IDs we want.
				set := []byte(`{
          "command": "message",
          "identifier": "{\"channel\":\"CGSimplePrice\"}",
          "data": "{\"coin_id\":[\"ethereum\"],\"action\":\"set_tokens\"}"
        }`)

				err = c.Write(set)
				if err != nil {
					tracer.Panic(err)
				}
			},
		})
	}

	go func() {
		err := str.Daemon()
		if err != nil {
			tracer.Panic(err)
		}
	}()

	go func() {
		{
			time.Sleep(90 * time.Second)
		}

		err := str.Close()
		if err != nil {
			tracer.Panic(err)
		}

		{
			close(dat)
		}
	}()

	for x := range dat {
		if bytes.Contains(x, []byte(`"c":"C1"`)) {
			fmt.Printf("%s\n", x)
		}
	}
}
