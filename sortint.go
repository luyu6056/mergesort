package mergesort

import "sync"

var intpool = sync.Pool{New: func() interface{} {
	return make([]int, 0)
}}

func SortInt(list []int) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] > list[i+2] {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] > list[i+3] {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, r_b, index, n int
	tmp := make([]int, max_len)
	step = 4
	for step < max_len {
		n++
		step <<= 1
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
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
						tmp[index] = list[l]
						l++
					case l == r_b:
						tmp[index] = list[r]
						r++
					case list[l] > list[r]:
						tmp[index] = list[r]
						r++
					default:
						tmp[index] = list[l]
						l++
					}
					index++
				}

			}
		} else {
			for i := 0; i < max_len; i += step {
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
						list[index] = tmp[l]
						l++
					case l == r_b:
						list[index] = tmp[r]
						r++
					case tmp[l] > tmp[r]:
						list[index] = tmp[r]
						r++
					default:
						list[index] = tmp[l]
						l++
					}
					index++
				}

			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}

func SortIntDesc(list []int) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] < list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] < list[i+2] {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] < list[i+3] {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, r_b, index, n int
	tmp := make([]int, max_len)
	step = 4
	for step < max_len {
		n++
		step <<= 1
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
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
						tmp[index] = list[l]
						l++
					case l == r_b:
						tmp[index] = list[r]
						r++
					case list[l] < list[r]:
						tmp[index] = list[r]
						r++
					default:
						tmp[index] = list[l]
						l++
					}
					index++
				}

			}
		} else {
			for i := 0; i < max_len; i += step {
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
						list[index] = tmp[l]
						l++
					case l == r_b:
						list[index] = tmp[r]
						r++
					case tmp[l] < tmp[r]:
						list[index] = tmp[r]
						r++
					default:
						list[index] = tmp[l]
						l++
					}
					index++
				}

			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}
func SortInt8(list []int8) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] > list[i+2] {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] > list[i+3] {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, r_b, index, n int
	tmp := make([]int8, max_len)
	step = 4
	for step < max_len {
		n++
		step <<= 1
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
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
						tmp[index] = list[l]
						l++
					case l == r_b:
						tmp[index] = list[r]
						r++
					case list[l] > list[r]:
						tmp[index] = list[r]
						r++
					default:
						tmp[index] = list[l]
						l++
					}
					index++
				}

			}
		} else {
			for i := 0; i < max_len; i += step {
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
						list[index] = tmp[l]
						l++
					case l == r_b:
						list[index] = tmp[r]
						r++
					case tmp[l] > tmp[r]:
						list[index] = tmp[r]
						r++
					default:
						list[index] = tmp[l]
						l++
					}
					index++
				}

			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}

func SortIntDesc8(list []int8) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] < list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] < list[i+2] {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] < list[i+3] {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, r_b, index, n int
	tmp := make([]int8, max_len)
	step = 4
	for step < max_len {
		n++
		step <<= 1
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
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
						tmp[index] = list[l]
						l++
					case l == r_b:
						tmp[index] = list[r]
						r++
					case list[l] < list[r]:
						tmp[index] = list[r]
						r++
					default:
						tmp[index] = list[l]
						l++
					}
					index++
				}

			}
		} else {
			for i := 0; i < max_len; i += step {
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
						list[index] = tmp[l]
						l++
					case l == r_b:
						list[index] = tmp[r]
						r++
					case tmp[l] < tmp[r]:
						list[index] = tmp[r]
						r++
					default:
						list[index] = tmp[l]
						l++
					}
					index++
				}

			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}
func SortInt16(list []int16) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] > list[i+2] {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] > list[i+3] {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, r_b, index, n int
	tmp := make([]int16, max_len)
	step = 4
	for step < max_len {
		n++
		step <<= 1
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
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
						tmp[index] = list[l]
						l++
					case l == r_b:
						tmp[index] = list[r]
						r++
					case list[l] > list[r]:
						tmp[index] = list[r]
						r++
					default:
						tmp[index] = list[l]
						l++
					}
					index++
				}

			}
		} else {
			for i := 0; i < max_len; i += step {
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
						list[index] = tmp[l]
						l++
					case l == r_b:
						list[index] = tmp[r]
						r++
					case tmp[l] > tmp[r]:
						list[index] = tmp[r]
						r++
					default:
						list[index] = tmp[l]
						l++
					}
					index++
				}

			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}

func SortIntDesc16(list []int16) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] < list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] < list[i+2] {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] < list[i+3] {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, r_b, index, n int
	tmp := make([]int16, max_len)
	step = 4
	for step < max_len {
		n++
		step <<= 1
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
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
						tmp[index] = list[l]
						l++
					case l == r_b:
						tmp[index] = list[r]
						r++
					case list[l] < list[r]:
						tmp[index] = list[r]
						r++
					default:
						tmp[index] = list[l]
						l++
					}
					index++
				}

			}
		} else {
			for i := 0; i < max_len; i += step {
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
						list[index] = tmp[l]
						l++
					case l == r_b:
						list[index] = tmp[r]
						r++
					case tmp[l] < tmp[r]:
						list[index] = tmp[r]
						r++
					default:
						list[index] = tmp[l]
						l++
					}
					index++
				}

			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}
func SortInt32(list []int32) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] > list[i+2] {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] > list[i+3] {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, r_b, index, n int
	tmp := make([]int32, max_len)
	step = 4
	for step < max_len {
		n++
		step <<= 1
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
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
						tmp[index] = list[l]
						l++
					case l == r_b:
						tmp[index] = list[r]
						r++
					case list[l] > list[r]:
						tmp[index] = list[r]
						r++
					default:
						tmp[index] = list[l]
						l++
					}
					index++
				}

			}
		} else {
			for i := 0; i < max_len; i += step {
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
						list[index] = tmp[l]
						l++
					case l == r_b:
						list[index] = tmp[r]
						r++
					case tmp[l] > tmp[r]:
						list[index] = tmp[r]
						r++
					default:
						list[index] = tmp[l]
						l++
					}
					index++
				}

			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}

func SortIntDesc32(list []int32) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] < list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] < list[i+2] {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] < list[i+3] {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, r_b, index, n int
	tmp := make([]int32, max_len)
	step = 4
	for step < max_len {
		n++
		step <<= 1
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
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
						tmp[index] = list[l]
						l++
					case l == r_b:
						tmp[index] = list[r]
						r++
					case list[l] < list[r]:
						tmp[index] = list[r]
						r++
					default:
						tmp[index] = list[l]
						l++
					}
					index++
				}

			}
		} else {
			for i := 0; i < max_len; i += step {
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
						list[index] = tmp[l]
						l++
					case l == r_b:
						list[index] = tmp[r]
						r++
					case tmp[l] < tmp[r]:
						list[index] = tmp[r]
						r++
					default:
						list[index] = tmp[l]
						l++
					}
					index++
				}

			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}
func SortInt64(list []int64) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] > list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] > list[i+2] {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] > list[i+3] {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] > list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, r_b, index, n int
	tmp := make([]int64, max_len)
	step = 4
	for step < max_len {
		n++
		step <<= 1
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
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
						tmp[index] = list[l]
						l++
					case l == r_b:
						tmp[index] = list[r]
						r++
					case list[l] > list[r]:
						tmp[index] = list[r]
						r++
					default:
						tmp[index] = list[l]
						l++
					}
					index++
				}

			}
		} else {
			for i := 0; i < max_len; i += step {
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
						list[index] = tmp[l]
						l++
					case l == r_b:
						list[index] = tmp[r]
						r++
					case tmp[l] > tmp[r]:
						list[index] = tmp[r]
						r++
					default:
						list[index] = tmp[l]
						l++
					}
					index++
				}

			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}

func SortInt64Desc(list []int64) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] < list[i+1] {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] < list[i+2] {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] < list[i+3] {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}

	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] < list[i+2] {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, r_b, index, n int
	tmp := make([]int64, max_len)
	step = 4
	for step < max_len {
		n++
		step <<= 1
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
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
						tmp[index] = list[l]
						l++
					case l == r_b:
						tmp[index] = list[r]
						r++
					case list[l] < list[r]:
						tmp[index] = list[r]
						r++
					default:
						tmp[index] = list[l]
						l++
					}
					index++
				}

			}
		} else {
			for i := 0; i < max_len; i += step {
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
						list[index] = tmp[l]
						l++
					case l == r_b:
						list[index] = tmp[r]
						r++
					case tmp[l] < tmp[r]:
						list[index] = tmp[r]
						r++
					default:
						list[index] = tmp[l]
						l++
					}
					index++
				}

			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}
