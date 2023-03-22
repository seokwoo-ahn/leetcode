func nextPermutation(nums []int)  {
    index := -1
    for i := len(nums) - 1; i > 0; i-- {
        if nums[i] > nums[i-1] {
            index = i-1
            break
        }
    }
    
    if index == -1 {
        sort.Ints(nums)
    } else {
        val := nums[index]
        sort.Ints(nums[index:len(nums)])
        
        for i := index; i < len(nums) - 1; i++ {
            if nums[i+1] > val {
                temp := nums[index]
                nums[index] = nums[i+1]
                nums[i+1] = temp
                break
            }
        }        
        sort.Ints(nums[index+1:])
    }
}
