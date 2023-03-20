func fourSum(nums []int, target int) [][]int {
    res := make([][]int, 0)
    sort.Ints(nums)
    for i := 0; i < len(nums) - 3; i++ {
        if i != 0 && nums[i] == nums[i-1]{
            continue
        }
        for j := i+1; j < len(nums) - 2; j++ {
            if j != i+1 && nums[j] == nums[j-1]{
                continue
            }
            k,l := j+1, len(nums) -1
            for k < l {
                sum := nums[i] + nums[j] + nums[k] + nums[l]
                if sum == target {
                    ans := []int{nums[i],nums[j],nums[k],nums[l]}
                    res = append(res, ans)
                    k++
                    l--
                    for k < l && nums[k] == nums[k-1]{
                        k++
                    }
                    for k < l && nums[l] == nums[l+1]{
                        l--
                    }
                } else if sum > target {
                    l--
                } else {
                    k++
                }
            }
        }
    }
    return res
}
