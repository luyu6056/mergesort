package mergesort

func SortInt(list []int) {
	var step, l, max, r, r_b, index, n int
	max_len := len(list)
	tmp := make([]int, max_len)

	step = 1
	for {
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
		if step > max_len {
			if n&1 == 1 {
				copy(list, tmp)
			}
			return
		}
	}
}

func SortIntDesc(list []int) {
	var step, l, max, r, r_b, index, n int
	max_len := len(list)
	tmp := make([]int, max_len)

	step = 1
	for {
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
		if step > max_len {
			if n&1 == 1 {
				copy(list, tmp)
			}
			return
		}
	}
}
func SortInt8(list []int8) {
	var step, l, max, r, r_b, index, n int
	max_len := len(list)
	tmp := make([]int8, max_len)

	step = 1
	for {
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
		if step > max_len {
			if n&1 == 1 {
				copy(list, tmp)
			}
			return
		}
	}
}
func SortInt8Desc(list []int8) {
	var step, l, max, r, r_b, index, n int
	max_len := len(list)
	tmp := make([]int8, max_len)

	step = 1
	for {
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
		if step > max_len {
			if n&1 == 1 {
				copy(list, tmp)
			}
			return
		}
	}
}
func SortInt16(list []int16) {
	var step, l, max, r, r_b, index, n int
	max_len := len(list)
	tmp := make([]int16, max_len)

	step = 1
	for {
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
		if step > max_len {
			if n&1 == 1 {
				copy(list, tmp)
			}
			return
		}
	}
}
func SortInt16Desc(list []int16) {
	var step, l, max, r, r_b, index, n int
	max_len := len(list)
	tmp := make([]int16, max_len)

	step = 1
	for {
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
		if step > max_len {
			if n&1 == 1 {
				copy(list, tmp)
			}
			return
		}
	}
}
func SortInt32(list []int32) {
	var step, l, max, r, r_b, index, n int
	max_len := len(list)
	tmp := make([]int32, max_len)

	step = 1
	for {
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
		if step > max_len {
			if n&1 == 1 {
				copy(list, tmp)
			}
			return
		}
	}
}
func SortInt32Desc(list []int32) {
	var step, l, max, r, r_b, index, n int
	max_len := len(list)
	tmp := make([]int32, max_len)

	step = 1
	for {
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
		if step > max_len {
			if n&1 == 1 {
				copy(list, tmp)
			}
			return
		}
	}
}
func SortInt64(list []int64) {
	var step, l, max, r, r_b, index, n int
	max_len := len(list)
	tmp := make([]int64, max_len)

	step = 1
	for {
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
		if step > max_len {
			if n&1 == 1 {
				copy(list, tmp)
			}
			return
		}
	}
}
func SortInt64Desc(list []int64) {
	var step, l, max, r, r_b, index, n int
	max_len := len(list)
	tmp := make([]int64, max_len)

	step = 1
	for {
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
		if step > max_len {
			if n&1 == 1 {
				copy(list, tmp)
			}
			return
		}
	}
}
