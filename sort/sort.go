package sort

// BubbleSort 冒泡排序
func BubbleSort(data []int) {
	n := len(data)
	if n == 0 {
		return
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

// SelectSort 选择排序
func SelectSort(data []int) {
	n := len(data)
	if n == 0 {
		return
	}

	for i := 0; i < n-1; i++ {
		idx := i
		min := data[i]
		for j := i + 1; j < n; j++ {
			if min > data[j] {
				idx = j
				min = data[j]
			}
		}
		if idx != i {
			data[i], data[idx] = data[idx], data[i]
		}
	}
}

// QuickSort 快速排序
func QuickSort(data []int) {
	n := len(data)
	if n == 0 {
		return
	}

	i, j := 0, n-1
	pivot := data[0]
	for i != j {
		for ; j > i; j-- {
			if data[j] < pivot {
				break
			}
		}

		for ; i < j; i++ {
			if data[i] > pivot {
				break
			}
		}

		if i < j {
			data[i], data[j] = data[j], data[i]
		}
	}

	data[0], data[i] = data[i], data[0]
	QuickSort(data[:i])
	QuickSort(data[i+1:])
}

// InsertSortV1 插入排序
func InsertSortV1(data []int) {
	n := len(data)
	if n == 0 {
		return
	}

	tmp := make([]int, n)
	tmp[0] = data[0]
	for i := 1; i < n; i++ {
		// fmt.Println(tmp)
		if data[i] >= tmp[i-1] {
			tmp[i] = data[i]
		} else {
			idx := 0
			for j := i - 2; j >= 0; j-- {
				if data[i] >= tmp[j] {
					idx = j + 1
					break
				}
			}
			// fmt.Println(idx)
			for k := i; k > idx; k-- {
				tmp[k] = tmp[k-1]
			}
			tmp[idx] = data[i]
		}
	}

	for i, v := range tmp {
		data[i] = v
	}
}

// InsertSortV2 插入排序
func InsertSortV2(data []int) {
	n := len(data)
	if n == 0 {
		return
	}

	current := 0
	for i := 0; i < n-1; i++ {
		current = data[i+1]
		preIndex := i
		for preIndex >= 0 && current < data[preIndex] {
			data[preIndex+1] = data[preIndex]
			preIndex--
		}
		data[preIndex+1] = current
	}
}

// ShellSort 希尔排序
func ShellSort(data []int) {
	n := len(data)
	if n == 0 {
		return
	}

	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			current := data[i]
			preIndex := i - gap
			for preIndex >= 0 && current < data[preIndex] {
				data[preIndex+gap] = data[preIndex]
				preIndex -= gap
			}
			data[preIndex+gap] = current
		}
		gap = gap / 2
	}
}

// MergeSort 归并排序
func MergeSort(data []int) []int {
	n := len(data)
	if n < 2 {
		return data
	}

	mid := n / 2

	return merge(MergeSort(data[:mid]), MergeSort(data[mid:]))
}

func merge(left, right []int) []int {
	leftLength := len(left)
	rightLength := len(right)
	resultLength := leftLength + rightLength
	result := make([]int, resultLength)
	for i, j, index := 0, 0, 0; index < resultLength; index++ {
		if i >= leftLength {
			result[index] = right[j]
			j++
		} else if j >= rightLength {
			result[index] = left[i]
			i++
		} else if left[i] > right[j] {
			result[index] = right[j]
			j++
		} else {
			result[index] = left[i]
			i++
		}
	}

	return result
}

func merge2(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	// Either left or right may have elements left; consume them.
	// (Only one of the following loops will actually be entered.)
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
