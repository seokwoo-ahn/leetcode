func searchInsert(nums []int, target int) int {
    start := 0
    end := len(nums) - 1
    
    for start <= end {
        mid := (start + end)/2
        
        if nums[mid] < target {
            start = mid + 1
        } else if nums[mid] > target {
            end = mid - 1
        } else {
            return mid
        }
    }
    return start
}
