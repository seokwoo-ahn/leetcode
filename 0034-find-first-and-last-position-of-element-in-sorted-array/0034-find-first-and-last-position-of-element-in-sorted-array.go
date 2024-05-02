func searchRange(nums []int, target int) []int {
    start, end := 0, len(nums) - 1
    
    for start <= end {
        mid := (start + end) / 2

        if nums[mid] == target {
            leftIdx := findIdx(nums, target, start, mid, true)
            rightIdx := findIdx(nums, target, mid, end, false)
            return []int{leftIdx, rightIdx}
        }
        
        if nums[mid] > target {
            end = mid - 1
        } else {
            start = mid + 1
        }
    }
    return []int{-1,-1}
}

func findIdx(nums []int, target int, start int, end int, isLeft bool) int {
    var res int
    if isLeft {
        res = end
    } else {
        res = start
    }

    for start <= end {
        mid := (start + end) / 2
        if nums[mid] == target {
            if isLeft {
                end = mid - 1
            } else {
                start = mid + 1
            }
            res = mid
        } else {
            if isLeft {
                start = mid + 1   
            } else {
                end = mid - 1         
            }
        }
    }
    return res
}
