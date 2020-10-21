package syncpooln_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/pratikdeoghare/syncpooln"
)

var v *[]byte

func BenchmarkSyncPoolGet(b *testing.B) {
	b.ReportAllocs()

	maxSize := 1000

	p := sync.Pool{
		New: func() interface{} {
			n := rand.Intn(maxSize)
			v := make([]byte, n)
			return &v
		},
	}

	for i := 0; i < b.N; i++ {
		v = p.Get().(*[]byte)
		p.Put(v)
	}
}

func BenchmarkPoolnGet(b *testing.B) {
	b.ReportAllocs()

	maxSize := 1000

	p := syncpooln.New(
		func(n int) interface{} {
			v := make([]byte, n)
			return &v
		},
	)

	for i := 0; i < b.N; i++ {
		n := rand.Intn(maxSize)
		v = p.Get(n).(*[]byte)
		p.Put(n, v)
	}

}

var mv map[string]string

func BenchmarkSyncPoolGetMaps(b *testing.B) {
	maxSize := 1000

	mv = make(map[string]string, maxSize)
	b.ReportAllocs()

	p := sync.Pool{
		New: func() interface{} {
			n := rand.Intn(maxSize)
			v := make(map[string]string, n)
			return v
		},
	}

	for i := 0; i < b.N; i++ {
		mv = p.Get().(map[string]string)
		p.Put(mv)
	}
}

func BenchmarkPoolnGetMaps(b *testing.B) {
	maxSize := 1000
	mv = make(map[string]string, maxSize)

	b.ReportAllocs()

	p := syncpooln.New(
		func(n int) interface{} {
			v := make(map[string]string, n)
			return v
		},
	)

	for i := 0; i < b.N; i++ {
		n := rand.Intn(maxSize)
		mv = p.Get(n).(map[string]string)
		p.Put(n, mv)
	}

}
