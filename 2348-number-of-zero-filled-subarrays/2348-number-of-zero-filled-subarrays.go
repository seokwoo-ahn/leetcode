func zeroFilledSubarray(nums []int) int64 {
    res := int64(0)
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            count++
        }
        if nums[i] != 0 && count != 0 {
            res += countSubArray(count)
            count = 0
        }
    }
    if count != 0{
        res += countSubArray(count)
    }
    return res
}

func countSubArray(count int) int64 {
    res := count*(count + 1) / 2
    return int64(res)
}