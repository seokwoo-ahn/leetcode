func threeSum(nums []int) [][]int {
    res := make([][]int, 0)
    sort.Ints(nums)
    for i := 0; i < len(nums) - 2; i++ {
        if nums[i] > 0 {
            return res
        }
        if i != 0 && nums[i] == nums[i-1]{
            continue   
        }
        j := i+1
        k := len(nums) - 1
        for j < k {
            sum := nums[j] + nums[k]
            if nums[i] + sum == 0{
                ans := []int{nums[i],nums[j],nums[k]}
                res = append(res, ans)
                j++
                k--
                for j < k && nums[j] == nums[j-1]{
                    j++
                }
                for j < k && nums[k] == nums[k+1]{
                    k--
                }
            } else if nums[i] + sum > 0 {
                k--
            } else if nums[i] + sum < 0{
                j++
            }
        }
    }
    return res
}
