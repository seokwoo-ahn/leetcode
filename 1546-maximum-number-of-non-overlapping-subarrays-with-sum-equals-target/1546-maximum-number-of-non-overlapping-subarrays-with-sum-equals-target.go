func maxNonOverlapping(nums []int, target int) int {
    start, res, sum := 0, 0, 0
    
    for i := 0; i < len(nums); i++ {
        sum = sum + nums[i]
        
        if sum == target {
            res++
            sum = 0
            start = i + 1
            continue
        }
        
        tempSum := sum
        tempStart := start
        for tempStart < i {
            tempSum = tempSum - nums[tempStart]
            if tempSum == target {
                res++
                sum = 0
                start = i + 1
                break
            }
            tempStart++
        }
    }
    return res
}
