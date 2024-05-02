func permute(nums []int) [][]int {
    res := make([][]int, 0)
    
    findAllPermutations(nums, make([]int, 0), &res)
    return res
}

func findAllPermutations(nums []int, val []int, res *[][]int) {
    if len(nums) == 0 {
        *res = append(*res, val)
        return
    }
    
    for i := 0; i < len(nums); i++ {
        tempVal := make([]int, len(val))
        copy(tempVal[:], val[:])
        tempVal = append(tempVal, nums[i])
        
        tempNums := make([]int, len(nums))
        copy(tempNums[:], nums[:])
        
        tempNums[i], tempNums[len(tempNums) - 1] = tempNums[len(tempNums) - 1], tempNums[i]
        tempNums = tempNums[:len(tempNums) - 1]
        findAllPermutations(tempNums, tempVal, res)
    }
}
