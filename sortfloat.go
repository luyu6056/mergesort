package mergesort

func isNaN32(f float32) bool {
	return f != f
}
func isNaN64(f float64) bool {
	return f != f
}
func SortFloat32(list []float32) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] > list[i+1] || !isNaN32(list[i]) && isNaN32(list[i+1]) {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] > list[i+2] || !isNaN32(list[i]) && isNaN32(list[i+2]) {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] > list[i+3] || !isNaN32(list[i+1]) && isNaN32(list[i+3]) {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] > list[i+2] || !isNaN32(list[i+1]) && isNaN32(list[i+2]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] > list[i+2] || !isNaN32(list[i]) && isNaN32(list[i+2]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] > list[i+2] || !isNaN32(list[i+1]) && isNaN32(list[i+2]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, index, n int
	tmp := make([]float32, max_len)
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
					if l == step/2+i || (r < max && (list[l] > list[r] || !isNaN32(list[l]) && isNaN32(list[r]))) {
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
					if l == step/2+i || (r < max && (tmp[l] > tmp[r] || !isNaN32(tmp[l]) && isNaN32(tmp[r]))) {
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

func SortFloat32Desc(list []float32) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] < list[i+1] || isNaN32(list[i]) && !isNaN32(list[i+1]) {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] < list[i+2] || isNaN32(list[i]) && !isNaN32(list[i+2]) {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] < list[i+3] || isNaN32(list[i+1]) && !isNaN32(list[i+3]) {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] < list[i+2] || isNaN32(list[i+1]) && !isNaN32(list[i+2]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] < list[i+2] || isNaN32(list[i]) && !isNaN32(list[i+2]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] < list[i+2] || isNaN32(list[i+1]) && !isNaN32(list[i+2]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, index, n int
	tmp := make([]float32, max_len)
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
					if l == step/2+i || (r < max && (list[l] < list[r] || isNaN32(list[l]) && !isNaN32(list[r]))) {
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
					if l == step/2+i || (r < max && (tmp[l] < tmp[r] || isNaN32(tmp[l]) && !isNaN32(tmp[r]))) {
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
func SortFloat64(list []float64) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] > list[i+1] || !isNaN64(list[i]) && isNaN64(list[i+1]) {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] > list[i+2] || !isNaN64(list[i]) && isNaN64(list[i+2]) {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] > list[i+3] || !isNaN64(list[i+1]) && isNaN64(list[i+3]) {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] > list[i+2] || !isNaN64(list[i+1]) && isNaN64(list[i+2]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] > list[i+2] || !isNaN64(list[i]) && isNaN64(list[i+2]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] > list[i+2] || !isNaN64(list[i+1]) && isNaN64(list[i+2]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, index, n int
	tmp := make([]float64, max_len)
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
					if l == step/2+i || (r < max && (list[l] > list[r] || !isNaN64(list[l]) && isNaN64(list[r]))) {
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
					if l == step/2+i || (r < max && (tmp[l] > tmp[r] || !isNaN64(tmp[l]) && isNaN64(tmp[r]))) {
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

func SortFloat64Desc(list []float64) {
	max_len := len(list)
	for i := 0; i < max_len-max_len&1; i += 2 {
		if list[i] < list[i+1] || isNaN64(list[i]) && !isNaN64(list[i+1]) {
			list[i], list[i+1] = list[i+1], list[i]
		}

	}
	for i := 0; i < max_len-max_len&3; i += 4 {
		if list[i] < list[i+2] || isNaN64(list[i]) && !isNaN64(list[i+2]) {
			list[i], list[i+2] = list[i+2], list[i]
		}
		if list[i+1] < list[i+3] || isNaN64(list[i+1]) && !isNaN64(list[i+3]) {
			list[i+1], list[i+3] = list[i+3], list[i+1]
		}
		if list[i+1] < list[i+2] || isNaN64(list[i+1]) && !isNaN64(list[i+2]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	if max_len&3 == 3 {
		i := max_len - 3
		if list[i] < list[i+2] || isNaN64(list[i]) && !isNaN64(list[i+2]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
			list[i], list[i+1] = list[i+1], list[i]
		} else if list[i+1] < list[i+2] || isNaN64(list[i+1]) && !isNaN64(list[i+2]) {
			list[i+1], list[i+2] = list[i+2], list[i+1]
		}
	}
	var step, l, max, r, index, n int
	tmp := make([]float64, max_len)
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
					if l == step/2+i || (r < max && (list[l] < list[r] || isNaN64(list[l]) && !isNaN64(list[r]))) {
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
					if l == step/2+i || (r < max && (tmp[l] < tmp[r] || isNaN64(tmp[l]) && !isNaN64(tmp[r]))) {
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
