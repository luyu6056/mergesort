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
var data_srot []int

func init() {
	rand.Seed(time.Now().UnixNano())
	data_int = make([]int, 9999)
	for k := range data_int {
		data_int[k] = rand.Int()
	}
	data_srot = make([]int, 9999)
	for k := range data_srot {
		data_srot[k] = k
	}

}

type list []int

func (l list) Len() int {
	return len(l)
}
func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l list) Less(i, j int) bool {
	return l[i] < l[j]
}

func Benchmark_sort_Ints_rand(b *testing.B) {
	b.StopTimer()
	c := make([]int, len(data_int))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(c, data_int)
		sort.Ints(c)
	}
}

func Benchmark_sort_Stable_rand(b *testing.B) {
	b.StopTimer()
	var c list = make([]int, len(data_int))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(c, data_int)
		sort.Stable(c)
	}

}

func Benchmark_mergesort_SortInt_rand(b *testing.B) {
	b.StopTimer()
	c := make([]int, len(data_int))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(c, data_int)
		mergesort.SortInt(c)
	}
}
func Benchmark_sort_Ints_isSorted(b *testing.B) {
	b.StopTimer()
	c := make([]int, len(data_srot))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(c, data_srot)
		sort.Ints(c)
	}
}

func Benchmark_sort_Stable_isSorted(b *testing.B) {
	b.StopTimer()
	var c list = make([]int, len(data_srot))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(c, data_srot)
		sort.Stable(c)
	}

}
func Benchmark_mergesort_SortInt_isSorted(b *testing.B) {
	b.StopTimer()
	c := make([]int, len(data_srot))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(c, data_srot)
		mergesort.SortInt(c)
	}
}
func Benchmark_sortAny(b *testing.B) {
	b.StopTimer()
	var c = make([]int, len(data_int))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(c, data_int)
		mergesort.SortAny(c, func(i, j unsafe.Pointer) bool {
			return *(*int)(i) < *(*int)(j)
		})
	}

}
