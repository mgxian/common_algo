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
