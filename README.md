# syncpooln 

`syncpooln` maintains a map of `int`s to `sync.Pool`s. This allows you to keep your objects organized into separate `sync.Pool`s by any criteria you want. It is safe for use by multiple goroutines simultaneously. 

#### Examples

##### Simple:

```go
package main

import (
	"bytes"
	"io"
	"os"

	"github.com/pratikdeoghare/syncpooln"
)

var bufPooln = syncpooln.New(
	func(n int) interface{} {
		return bytes.NewBuffer(make([]byte, n))
	})

func Log(w io.Writer, text string) {
	n := len(text)
	b := bufPooln.Get(n).(*bytes.Buffer)
	b.Reset()
	b.WriteString(text)
	w.Write(b.Bytes())
	bufPooln.Put(n, b)
}

func main() {
	Log(os.Stdout, "small\n")
	Log(os.Stdout, "and then he typed a very long string of characters to test this\n")
	// Output:
	// small
	// and then he typed a very long string of characters to test this
}
```

##### Little crazy: 

Only try this at home. This is just to show what can be done. Use caution. 

```go
package main 

import( 
	"github.com/pratikdeoghare/syncpooln"
)

func main() {
	mixPool := syncpooln.New(func(n int) interface{} {
		if n == 0 {
			return make([]int, 0)
		}
		return make(map[int]struct{})
	})

	// n == 0 -> a list type
	a := mixPool.Get(0).([]int)
	mixPool.Put(0, a)

	// n != 0 -> a set type
	b := mixPool.Get(1).(map[int]struct{})
	mixPool.Put(1, b)

	// Output:
}
```

#### References
1. https://golang.org/pkg/sync/#Poolo
2. https://en.wikipedia.org/wiki/Equivalence_class
