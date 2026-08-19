func findMaxConsecutiveOnes(nums []int) int {
    var count, max int
    for _, num := range nums {
        if num == 1 {
            count++
            if max < count {
                max = count
            }
        } else {
            count = 0
        }
    }
    return max
}
