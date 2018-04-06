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

// InsertSort 插入排序
func InsertSort(data []int) {
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
