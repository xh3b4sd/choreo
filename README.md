# choreo

This package is a collection of simple execution path primitives like retrying,
timing out and runing business logic concurrently.

### Backoff

```golang
package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/xh3b4sd/choreo/backoff"
)

func main() {
	var tok *backoff.Token
	{
		tok = backoff.New(backoff.Config{
			Bac: []time.Duration{
				1 * time.Second,
				2 * time.Second,
				3 * time.Second,
			},
		})
	}

	fnc := func() error {
		fmt.Printf("tried at second %d\n", time.Now().Second())
		return errors.New("error")
	}

	err := tok.Backoff(fnc)
	if err != nil {
		fmt.Printf("%s at second %d\n", err.Error(), time.Now().Second())
	}
}
```

```
go run ./backoff/example/
```

```
tried at second 0
tried at second 1
tried at second 3
error at second 6
```

### Jitter

```golang
package main

import (
	"fmt"
	"time"

	"github.com/xh3b4sd/choreo/jitter"
)

func main() {
	jit := jitter.New[time.Duration](jitter.Config{
		Per: 0.1,
	})

	for range 5 {
		fmt.Printf("%s\n", jit.Percent(time.Minute))
	}
}
```

```
go run ./jitter/example/
```

```
55.771641992s
1m4.528653402s
59.012041686s
55.326155258s
1m3.387997178s
```

### Parallel

```golang
package main

import (
	"fmt"

	"github.com/xh3b4sd/choreo/parallel"
)

func main() {
	var inp []string
	{
		inp = []string{"10", "20", "30", "40", "50"}
	}

	fnc := func(i int, x string) error {
		fmt.Printf("index %d value %s\n", i, x)
		return nil
	}

	err := parallel.Slice(inp, fnc)
	if err != nil {
		panic(err)
	}
}
```

```
go run ./parallel/example/
```

```
// index 1 value 20
// index 2 value 30
// index 4 value 50
// index 0 value 10
// index 3 value 40
```

### Success

```golang
package main

import (
	"fmt"

	"github.com/xh3b4sd/choreo/success"
	"golang.org/x/sync/errgroup"
)

func main() {
	var mut *success.Mutex
	{
		mut = success.New(success.Config{
			Suc: 2,
		})
	}

	var grp errgroup.Group
	{
		grp = errgroup.Group{}
	}

	for range 100 {
		grp.Go(func() error {
			return mut.Success(func() error {
				fmt.Printf("hello world\n")
				return nil
			})
		})
	}

	err := grp.Wait()
	if err != nil {
		panic(err)
	}
}
```

```
go run ./success/example/
```

```
hello world
hello world
```
