func maxSubArray(nums []int) int {
    return findMaxSubArraySum(nums, 0, len(nums)-1)
}

func findMaxCenterSubArraySum(nums []int, start, end, mid int) (int){
    maxLeftSum := -1 << 31
    leftSum := 0
    for i := mid; i >= start; i-- {
        leftSum += nums[i]
        if maxLeftSum < leftSum {
            maxLeftSum = leftSum
        }
    }
    
    maxRightSum := -1 << 31
    rightSum := 0
    for i := mid+1; i <= end; i++ {
        rightSum += nums[i]
        if maxRightSum < rightSum {
            maxRightSum = rightSum
        }
    }
    
    return maxLeftSum + maxRightSum
}

func findMaxSubArraySum(nums []int, start, end int) (int) {
    if start == end {
        return nums[start]
    }
    
    mid := (start + end) / 2
    maxLeftSum := findMaxSubArraySum(nums, start, mid)
    maxRightSum := findMaxSubArraySum(nums, mid+1, end)
    maxCenterSum := findMaxCenterSubArraySum(nums, start, end, mid)
    
    res := maxLeftSum
    if res < maxRightSum {
        res = maxRightSum
    }
    
    if res < maxCenterSum {
        res = maxCenterSum
    }
    
    return res
}