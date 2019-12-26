package mergesort

import (
	"errors"
	"reflect"
	"unsafe"
)

type SliceHeader struct {
	Dataptr uintptr
	Len     uintptr
	Cap     int
}

func SortAny(list interface{}, order func(small, big unsafe.Pointer) bool) error {
	t := reflect.TypeOf(list)
	if t.Kind() != reflect.Slice {
		return errors.New("not a slice")
	}
	truetype := t.Elem()
	size := truetype.Size()
	p := (*[2]unsafe.Pointer)(unsafe.Pointer(&list))[1]
	list_sh := (*SliceHeader)(p)
	var step, l, max, r, r_b, index, n uintptr
	max_len := list_sh.Len
	tmp := reflect.MakeSlice(t, int(max_len), int(max_len))
	list_ptr := list_sh.Dataptr
	tmp_ptr := tmp.Pointer()
	for i := uintptr(0); i < max_len-max_len&1; i += 2 {
		if order(unsafe.Pointer(list_ptr+i*size), unsafe.Pointer(list_ptr+(i+1)*size)) {
			swap(list_ptr+i*size, list_ptr+(i+1)*size, size)
		}

	}
	for i := uintptr(0); i < max_len-max_len&3; i += 4 {
		if order(unsafe.Pointer(list_ptr+i*size), unsafe.Pointer(list_ptr+(i+2)*size)) {
			swap(list_ptr+i*size, list_ptr+(i+2)*size, size)
		}
		if order(unsafe.Pointer(list_ptr+(i+1)*size), unsafe.Pointer(list_ptr+(i+3)*size)) {
			swap(list_ptr+(i+1)*size, list_ptr+(i+3)*size, size)
		}
		if order(unsafe.Pointer(list_ptr+(i+1)*size), unsafe.Pointer(list_ptr+(i+2)*size)) {
			swap(list_ptr+(i+1)*size, list_ptr+(i+2)*size, size)
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if order(unsafe.Pointer(list_ptr+i*size), unsafe.Pointer(list_ptr+(i+2)*size)) {
			swap(list_ptr+(i+1)*size, list_ptr+(i+2)*size, size)
			swap(list_ptr+(i)*size, list_ptr+(i+1)*size, size)
		} else if order(unsafe.Pointer(list_ptr+(i+1)*size), unsafe.Pointer(list_ptr+(i+2)*size)) {
			swap(list_ptr+(i+1)*size, list_ptr+(i+2)*size, size)
		}
	}
	step = 1
	for step < max_len {
		n++
		step <<= 1
		if n&1 == 1 {
			for i := uintptr(0); i < max_len; i += step {
				l = i
				max = i + step
				r = (max-l)/2 + l
				r_b = r
				if max > max_len {
					max = max_len
				}
				index = i
				for index < max {
					switch {
					case r >= max:
						unsafeset(tmp_ptr+index*size, list_ptr+l*size, size)
						l++
					case l == r_b:
						unsafeset(tmp_ptr+index*size, list_ptr+r*size, size)
						r++
					case order(unsafe.Pointer(list_ptr+l*size), unsafe.Pointer(list_ptr+r*size)):
						unsafeset(tmp_ptr+index*size, list_ptr+r*size, size)
						r++
					default:
						unsafeset(tmp_ptr+index*size, list_ptr+l*size, size)
						l++
					}
					index++
				}

			}
		} else {
			for i := uintptr(0); i < max_len; i += step {
				l = i
				max = i + step
				r = (max-l)/2 + l
				r_b = r
				if max > max_len {
					max = max_len
				}
				index = i
				for index < max {
					switch {
					case r >= max:
						unsafeset(list_ptr+index*size, tmp_ptr+l*size, size)
						l++
					case l == r_b:
						unsafeset(list_ptr+index*size, tmp_ptr+r*size, size)
						r++
					case order(unsafe.Pointer(tmp_ptr+l*size), unsafe.Pointer(tmp_ptr+r*size)):
						unsafeset(list_ptr+index*size, tmp_ptr+r*size, size)
						r++
					default:
						unsafeset(list_ptr+index*size, tmp_ptr+l*size, size)
						l++
					}
					index++
				}

			}
		}

	}
	if n&1 == 1 {
		reflect.Copy(reflect.ValueOf(list), tmp)
	}
	return nil
}
func unsafeset(dst, src, size uintptr) {
	for size > 0 {
		switch size & 7 {
		case 0:
			*(*uint64)(unsafe.Pointer(dst + size - 8)) = *(*uint64)(unsafe.Pointer(src + size - 8))
			size -= 8
		case 7, 6, 5, 4:
			*(*uint32)(unsafe.Pointer(dst + size - 4)) = *(*uint32)(unsafe.Pointer(src + size - 4))
			size -= 4
		case 3, 2:
			*(*uint16)(unsafe.Pointer(dst + size - 2)) = *(*uint16)(unsafe.Pointer(src + size - 2))
			size -= 2
		case 1:
			*(*uint8)(unsafe.Pointer(dst + size - 1)) = *(*uint8)(unsafe.Pointer(src + size - 1))
			size -= 1
		}
	}

}
func swap(dst, src, size uintptr) {
	for size > 0 {
		switch size & 7 {
		case 0:
			*(*uint64)(unsafe.Pointer(dst + size - 8)), *(*uint64)(unsafe.Pointer(src + size - 8)) = *(*uint64)(unsafe.Pointer(src + size - 8)), *(*uint64)(unsafe.Pointer(dst + size - 8))
			size -= 8
		case 7, 6, 5, 4:
			*(*uint64)(unsafe.Pointer(dst + size - 4)), *(*uint64)(unsafe.Pointer(src + size - 4)) = *(*uint64)(unsafe.Pointer(src + size - 4)), *(*uint64)(unsafe.Pointer(dst + size - 4))
			size -= 4
		case 3, 2:
			*(*uint64)(unsafe.Pointer(dst + size - 2)), *(*uint64)(unsafe.Pointer(src + size - 2)) = *(*uint64)(unsafe.Pointer(src + size - 2)), *(*uint64)(unsafe.Pointer(dst + size - 2))
			size -= 2
		case 1:
			*(*uint64)(unsafe.Pointer(dst + size - 1)), *(*uint64)(unsafe.Pointer(src + size - 1)) = *(*uint64)(unsafe.Pointer(src + size - 1)), *(*uint64)(unsafe.Pointer(dst + size - 1))
			size -= 1
		}
	}
}
