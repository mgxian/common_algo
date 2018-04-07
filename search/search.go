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
