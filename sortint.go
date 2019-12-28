package mergesort

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
	var step, l, max, r, index, n int
	tmp := make([]int, max_len)
	step = 4
	for step < max_len {
		step <<= 1
		n++
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && list[l] > list[r]) {
						tmp[index] = list[r]
						r++
					} else {
						tmp[index] = list[l]
						l++
					}
				}
			}
		} else {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && tmp[l] > tmp[r]) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
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
	var step, l, max, r, index, n int
	tmp := make([]int, max_len)
	step = 4
	for step < max_len {
		step <<= 1
		n++
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && list[l] < list[r]) {
						tmp[index] = list[r]
						r++
					} else {
						tmp[index] = list[l]
						l++
					}
				}
			}
		} else {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && tmp[l] < tmp[r]) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
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
	var step, l, max, r, index, n int
	tmp := make([]int8, max_len)
	step = 4
	for step < max_len {
		step <<= 1
		n++
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && list[l] > list[r]) {
						tmp[index] = list[r]
						r++
					} else {
						tmp[index] = list[l]
						l++
					}
				}
			}
		} else {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && tmp[l] > tmp[r]) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
				}
			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}

func SortInt8Desc(list []int8) {
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
	var step, l, max, r, index, n int
	tmp := make([]int8, max_len)
	step = 4
	for step < max_len {
		step <<= 1
		n++
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && list[l] < list[r]) {
						tmp[index] = list[r]
						r++
					} else {
						tmp[index] = list[l]
						l++
					}
				}
			}
		} else {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && tmp[l] < tmp[r]) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
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
	var step, l, max, r, index, n int
	tmp := make([]int16, max_len)
	step = 4
	for step < max_len {
		step <<= 1
		n++
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && list[l] > list[r]) {
						tmp[index] = list[r]
						r++
					} else {
						tmp[index] = list[l]
						l++
					}
				}
			}
		} else {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && tmp[l] > tmp[r]) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
				}
			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}

func SortInt16Desc(list []int16) {
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
	var step, l, max, r, index, n int
	tmp := make([]int16, max_len)
	step = 4
	for step < max_len {
		step <<= 1
		n++
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && list[l] < list[r]) {
						tmp[index] = list[r]
						r++
					} else {
						tmp[index] = list[l]
						l++
					}
				}
			}
		} else {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && tmp[l] < tmp[r]) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
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
	var step, l, max, r, index, n int
	tmp := make([]int32, max_len)
	step = 4
	for step < max_len {
		step <<= 1
		n++
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && list[l] > list[r]) {
						tmp[index] = list[r]
						r++
					} else {
						tmp[index] = list[l]
						l++
					}
				}
			}
		} else {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && tmp[l] > tmp[r]) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
				}
			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}

func SortInt32Desc(list []int32) {
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
	var step, l, max, r, index, n int
	tmp := make([]int32, max_len)
	step = 4
	for step < max_len {
		step <<= 1
		n++
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && list[l] < list[r]) {
						tmp[index] = list[r]
						r++
					} else {
						tmp[index] = list[l]
						l++
					}
				}
			}
		} else {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && tmp[l] < tmp[r]) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
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
	var step, l, max, r, index, n int
	tmp := make([]int64, max_len)
	step = 4
	for step < max_len {
		step <<= 1
		n++
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && list[l] > list[r]) {
						tmp[index] = list[r]
						r++
					} else {
						tmp[index] = list[l]
						l++
					}
				}
			}
		} else {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && tmp[l] > tmp[r]) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
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
	var step, l, max, r, index, n int
	tmp := make([]int64, max_len)
	step = 4
	for step < max_len {
		step <<= 1
		n++
		if n&1 == 1 {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && list[l] < list[r]) {
						tmp[index] = list[r]
						r++
					} else {
						tmp[index] = list[l]
						l++
					}
				}
			}
		} else {
			for i := 0; i < max_len; i += step {
				l = i
				max = i + step
				r = step/2 + l
				if max > max_len {
					max = max_len
				}
				for index = i; index < max; index++ {
					if l == step/2+i || (r < max && tmp[l] < tmp[r]) {
						list[index] = tmp[r]
						r++
					} else {
						list[index] = tmp[l]
						l++
					}
				}
			}
		}
	}
	if n&1 == 1 {
		copy(list, tmp)
	}
}
