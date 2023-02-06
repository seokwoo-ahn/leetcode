func shuffle(nums []int, n int) []int {
    var result []int
    
    for i:=0; i< n; i++ {
        result = append(result, nums[i], nums[i+n])
    }
    return result
}