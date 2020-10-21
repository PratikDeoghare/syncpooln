package syncpooln_test

import (
	"bytes"
	"fmt"
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
	fmt.Println("Buffer length: ", b.Len())
	b.Reset()
	b.WriteString(text)
	w.Write(b.Bytes())
	bufPooln.Put(n, b)
}

func ExamplePooln() {
	Log(os.Stdout, "small\n")
	Log(os.Stdout, "and then he typed a very long string of characters to test this\n")
	// Output:
	// Buffer length:  6
	// small
	// Buffer length:  64
	// and then he typed a very long string of characters to test this
}

func ExampleCrazy() {
	mixPool := syncpooln.New(func(n int) interface{} {
		if n == 0 {
			return make([]int, 0)
		}
		return make(map[int]struct{})
	})

	// == 0 -> a list type
	a := mixPool.Get(0).([]int)
	mixPool.Put(0, a)

	// != 0 -> a set type
	b := mixPool.Get(1).(map[int]struct{})
	mixPool.Put(1, b)

	// Output:
}
