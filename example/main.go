package main

import (
	"fmt"
	"unsafe"

	"github.com/luyu6056/mergesort"
)

var data_pointer []*Test
var data_struct []Test

type Test struct {
	Name   string
	Height float64
	Age    int
}

func main() {
	data_pointer = append(data_pointer, &Test{
		Name:   "one",
		Height: 1.72,
		Age:    20,
	})
	data_pointer = append(data_pointer, &Test{
		Name:   "two",
		Height: 1.80,
		Age:    19,
	})
	data_pointer = append(data_pointer, &Test{
		Name:   "three",
		Height: 1.72,
		Age:    18,
	})
	data_struct = append(data_struct, Test{
		Name:   "one",
		Height: 1.72,
		Age:    20,
	})
	data_struct = append(data_struct, Test{
		Name:   "two",
		Height: 1.80,
		Age:    19,
	})
	data_struct = append(data_struct, Test{
		Name:   "three",
		Height: 1.72,
		Age:    18,
	})
	SortAnyStructPointer()
	SortAnyStructPointer2()
	SortAnyStruct()
	SortAnyStruct2()
}
func SortAnyStructPointer() {

	mergesort.SortAny(data_pointer, func(small, big unsafe.Pointer) bool { //small表示值小，big表示值大
		s := *(**Test)(small) //*Test
		b := *(**Test)(big)   //*Test
		if s.Height == b.Height {
			//务必 return A>B 的格式
			return s.Age > b.Age //年龄small小的优先
		}
		return b.Height > s.Height //身高big，比较高的优先
	})
	for _, v := range data_pointer {
		fmt.Printf("%+v\r\n", v)
	}
	//&{Name:two Height:1.8 Age:19}
	//&{Name:three Height:1.72 Age:18}
	//&{Name:one Height:1.72 Age:20}
}
func SortAnyStructPointer2() {

	mergesort.SortAny(data_pointer, func(small, big unsafe.Pointer) bool { //small表示值小，big表示值大
		s := *(**Test)(small)
		b := *(**Test)(big)
		if b.Height > s.Height {
			return true //优先 return true
		}
		if b.Height < s.Height {
			return false
		}
		if s.Age > b.Age {
			return true
		}
		return false //最后return false
	})
	for _, v := range data_pointer {
		fmt.Printf("%+v\r\n", v)
	}
	//&{Name:two Height:1.8 Age:19}
	//&{Name:three Height:1.72 Age:18}
	//&{Name:one Height:1.72 Age:20}
}
func SortAnyStruct() {

	mergesort.SortAny(data_struct, func(small, big unsafe.Pointer) bool { //small表示值小，big表示值大
		s := *(*Test)(small) //Test
		b := *(*Test)(big)   //Test
		if s.Height == b.Height {
			//务必 return A>B 的格式
			return s.Age > b.Age //年龄small小的优先
		}
		return b.Height > s.Height //身高big，比较高的优先
	})
	for _, v := range data_struct {
		fmt.Printf("%+v\r\n", v)
	}
	//{Name:two Height:1.8 Age:19}
	//{Name:three Height:1.72 Age:18}
	//{Name:one Height:1.72 Age:20}
}
func SortAnyStruct2() {
	mergesort.SortAny(data_struct, func(small, big unsafe.Pointer) bool { //small表示值小，big表示值大
		s := *(*Test)(small)
		b := *(*Test)(big)
		if b.Height > s.Height {
			return true //优先 return true
		}
		if b.Height < s.Height {
			return false
		}
		if s.Age > b.Age {
			return true
		}
		return false //最后return false
	})
	for _, v := range data_struct {
		fmt.Printf("%+v\r\n", v)
	}
	//{Name:two Height:1.8 Age:19}
	//{Name:three Height:1.72 Age:18}
	//{Name:one Height:1.72 Age:20}
}
