# mergesort
对于int，float，这个是比标准库sort.Ints,sort.Float64s要快一倍。
对于任意切片类型排序，比标准库sort.Stable()要快的稳定归并排序，无需定义Less，Len，Swap方法，但是要传递一个Less的func，比sort.Sort()略慢。
sortAny是基于反射复制一个切片，使用传递过来的less进行比较，使用unsafe定位切片成员地址，以及复制切片成员值。
使用方法在example文件夹内
# Benchmark 随机rand 999个切片
这是一个标准的归并排序，所以空间复杂度为O(n)

| | ns/op | allocation bytes | allocation times |
| --- | --- | --- | --- |
| sort.Ints | 86489 ns/op | 32 B/op | 1 allocs/op |
| mergesort.SortInt | 43584 ns/op | 8192 B/op | 1 allocs/op |
| sort.Float64s | 95562 ns/op | 32 B/op | 1 allocs/op |
| mergesort.SortFloat64 | 55183 ns/op | 8192 B/op | 1 allocs/op |
| sort.Sort | 91584 ns/op | 32 B/op | 1 allocs/op |
| sort.Stable | 243404 ns/op | 32 B/op | 1 allocs/op |
| mergesort.SortAny | 130290 ns/op | 24640 B/op | 3 allocs/op |
| sort.Stable | 100288 ns/op | 24576 B/op | 1 allocs/op |

#Usage SrotAny(interface{}, func(small, big unsafe.Pointer) bool)

int asc
```go
"github.com/luyu6056/mergesort"
var s []int
mergesort.SortAny(s, func(small, big unsafe.Pointer) bool {
		return (*(*int)(small)).score > (*(*int)(big)).score 
})
```
int desc
```go
"github.com/luyu6056/mergesort"
var s []int
mergesort.SortAny(s, func(small, big unsafe.Pointer) bool {
		return (*(*int)(small)).score < (*(*int)(big)).score 
})
```
struct
```go
"github.com/luyu6056/mergesort"
type A struct{
  B int
}
var s []A
mergesort.SortAny(s, func(small, big unsafe.Pointer) bool {
		return (*(*A)(small)).B < (*(*A)(big)).B 
})
```
pointer *struct
```go
"github.com/luyu6056/mergesort"
type A struct{
  B int
}
var s []*A
mergesort.SortAny(s, func(small, big unsafe.Pointer) bool {
		return (*(**A)(small)).B < (*(**A)(big)).B 
})
```
