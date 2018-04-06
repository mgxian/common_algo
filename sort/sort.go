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
