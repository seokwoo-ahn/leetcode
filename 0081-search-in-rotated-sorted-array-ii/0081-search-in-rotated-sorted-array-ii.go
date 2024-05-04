func search(nums []int, target int) bool {
	start, end := 0, len(nums)-1

	return loop(nums, target, start, end)
}

func loop(nums []int, target, start, end int) bool {
	for start <= end {
		mid := (start + end) / 2

		if nums[mid] == target {
			return true
		}

		if nums[mid] < nums[end] {
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else if nums[mid] > nums[end] {
			if nums[mid] > target && target >= nums[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			return loop(nums, target, start, mid-1) || loop(nums, target, mid+1, end)
		}
	}
	return false
}
