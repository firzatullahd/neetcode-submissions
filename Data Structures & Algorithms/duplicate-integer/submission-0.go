func hasDuplicate(nums []int) bool {
    mapNum := make(map[int]int)
	for _,v := range nums {
		mapNum[v]++
		if mapNum[v] > 1 {
			return true
		}
	}

	return false
}
