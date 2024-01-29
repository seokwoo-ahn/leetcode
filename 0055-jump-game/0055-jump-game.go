func canJump(nums []int) bool {
    if len(nums) == 1 {
        return true
    }
    
    maxIndex := 0
    for i := 0; i < len(nums) - 1; i++ {
        if i + nums[i] > maxIndex {
            maxIndex = i + nums[i]
        }
        if maxIndex >= len(nums) - 1 {
            return true
        }
        if maxIndex <= i {
            return false
        }
    }
    return false
}