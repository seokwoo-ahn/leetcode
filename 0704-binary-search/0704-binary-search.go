func search(nums []int, target int) int {
    first := 0
    end := len(nums) - 1
    
    for first <= end {
        mid := (first + end)/2
        if nums[mid] == target{
            return mid
        }
        
        if nums[mid] > target {
            end = mid - 1
        } else {
            first = mid + 1
        }
    }
    return -1
}