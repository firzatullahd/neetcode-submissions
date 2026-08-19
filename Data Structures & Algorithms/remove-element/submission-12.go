func removeElement(nums []int, val int) int {
   newNums := make([]int,0)

   for _,v := range nums {
		if v != val {
			newNums = append(newNums,v)
		}
   }

	for i := range newNums {
		nums[i] = newNums[i]
	}

   return len(newNums)
}
