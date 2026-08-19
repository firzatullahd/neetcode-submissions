func removeElement(nums []int, val int) int {
  	i := 0
	n := len(nums)

	for i< n {
		if nums[i] == val {
			nums[i] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			n--
		} else {
			i++
		}
	}

  return n
}
