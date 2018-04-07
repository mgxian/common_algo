package search

// SequenceSearch 顺序查找
func SequenceSearch(data []int, value int) (index int, founded bool) {
	for idx, v := range data {
		if v == value {
			index = idx
			founded = true
			return
		}
	}

	return
}

// BinarySearchV1 二分查找
func BinarySearchV1(data []int, value int, low, high int) (index int, founded bool) {
	mid := low + (high-low)/2
	if data[mid] == value {
		index = mid
		founded = true
		return
	}

	if data[mid] > value {
		index, founded = BinarySearchV1(data, value, low, mid-1)
	} else {
		index, founded = BinarySearchV1(data, value, mid+1, high)
	}

	return
}

// BinarySearchV2 二分查找
func BinarySearchV2(data []int, value int) (index int, founded bool) {
	low, high := 0, len(data)-1
	for low <= high {
		mid := low + (high-low)/2
		if data[mid] == value {
			index = mid
			founded = true
			return
		}

		if data[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return
}

// InsertionSearch 插值查找
func InsertionSearch(data []int, value int) (index int, founded bool) {
	low, high := 0, len(data)-1
	for low <= high {
		// mid := low + 1/2*(high-low)
		mid := low + (value-data[low])/(data[high]-data[low])*(high-low)
		if data[mid] == value {
			index = mid
			founded = true
			return
		}

		if data[mid] > value {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return
}

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// FibonacciSearch 斐波那契查找
func FibonacciSearch(data []int, value int) (index int, founded bool) {
	low, high, k := 0, len(data)-1, 0

	for len(data) > fibonacci(k)-1 {
		k++
	}
	// fmt.Println("K -----> ", k)
	tmp := make([]int, fibonacci(k)-1)
	copy(tmp, data)
	for i := len(data); i < fibonacci(k)-1; i++ {
		tmp[i] = data[len(data)-1]
	}

	// fmt.Println(tmp)
	for low <= high {
		mid := low + fibonacci(k-1) - 1
		// fmt.Println("mid ----->", mid)
		if value < tmp[mid] {
			high = mid - 1
			k--
		} else if value > tmp[mid] {
			low = mid + 1
			k -= 2
		} else {
			if mid < len(data) {
				index = mid
				founded = true
			} else {
				index = len(data) - 1
				founded = true
			}
			return
		}
	}
	return
}
