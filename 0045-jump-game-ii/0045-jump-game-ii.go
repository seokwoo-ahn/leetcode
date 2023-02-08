func jump(nums []int) int {
    max := 0
    cur := 0
    jumps := 0
    
    for i:=0; i<len(nums)-1; i++ {
        max = int(math.Max(float64(max), float64(i+nums[i])))
        if cur == i {
            jumps++
            cur = max
        }
        
        if cur > len(nums) -1 {
            return jumps
        }
    }
    return jumps
}