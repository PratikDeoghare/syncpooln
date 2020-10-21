package syncpooln_test

import (
	"sync"
	"testing"

	"github.com/pratikdeoghare/syncpooln"
)

func TestPoolnConcurrent(t *testing.T) {
	p := syncpooln.New(
		func(n int) interface{} {
			return make([]string, n)
		},
	)

	n := 100

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			wg.Done()
			x := p.Get(i).([]string)
			p.Put(i, x)
		}(i)
	}

	wg.Wait()

}
