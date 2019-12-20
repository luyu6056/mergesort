package test

import (
	"math/rand"
	"sort"
	"testing"
	"time"
	"unsafe"

	"github.com/luyu6056/mergesort"
)

var data_int []int

func init() {
	rand.Seed(time.Now().UnixNano())
	data_int = make([]int, 9999)
	for k := range data_int {
		data_int[k] = rand.Int()
	}

}

func Benchmark_Stand(b *testing.B) {
	b.StopTimer()
	c := make([]int, len(data_int))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		copy(c, data_int)
		sort.Ints(c)

	}
}
func Benchmark_SortInt(b *testing.B) {
	b.StopTimer()
	c := make([]int, len(data_int))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(c, data_int)
		mergesort.SortInt(c)
	}
}
func Benchmark_SortAny(b *testing.B) {
	b.StopTimer()
	c := make([]int, len(data_int))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(c, data_int)
		mergesort.SortAny(c, func(small, big unsafe.Pointer) bool {
			return *(*int)(small) > *(*int)(big)
		})
	}
}
