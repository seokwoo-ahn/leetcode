func findMaxK(nums []int) int {
    sort.Ints(nums)
    
    for i := len(nums) - 1; i >= 0; i-- {
        start, end, target := 0, i, -nums[i]
        
        if target < nums[start] {
            continue
        }
        
        if nums[i] < 0 {
            return -1
        }
        
        for start <= end {
            mid := (start + end) / 2
            
            if nums[mid] == target {
                return nums[i]
            }
            
            if target < nums[mid] {
                end = mid - 1
            } else {
                start = mid + 1
            }
        }
    }
    
    return -1
}