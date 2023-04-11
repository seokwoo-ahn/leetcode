func sortArrayByParityII(nums []int) []int {
    oddIdx := make([]int, 0)
    evenIdx := make([]int, 0)
    
    for i := 0; i < len(nums); i++ {
        if i%2 == 0 && nums[i]%2 != 0 {
            evenIdx = append(evenIdx, i)
        } else if i%2 == 1 && nums[i]%2 != 1 {
            oddIdx = append(oddIdx, i)
        } else {
            continue
        }
    }
    
    for len(oddIdx) > 0 {
        odd := oddIdx[0]
        oddIdx = oddIdx[1:]
        even := evenIdx[0]
        evenIdx = evenIdx[1:]
        
        temp := nums[odd]
        nums[odd] = nums[even]
        nums[even] = temp
    }
    return nums
}
