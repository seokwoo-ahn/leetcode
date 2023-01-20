func findSubsequences(nums []int) [][]int {
    var sequence []int
    var result [][]int
    backtrack(nums, 0, sequence, &result)
    return result
}

func backtrack(nums []int, index int, sequence []int, result *[][]int){
    if index == len(nums){
        if len(sequence) > 1 {
			*result = append(*result, sequence)
        }
        return
    }
     
    if len(sequence) == 0 || sequence[len(sequence) - 1] <= nums[index] {
        newSequence := make([]int, len(sequence))
		copy(newSequence, sequence)
		backtrack(nums, index+1, append(newSequence, nums[index]), result)
    }
    
    if len(sequence) > 0 && sequence[len(sequence) - 1] == nums[index] {
        return
    }
        backtrack(nums, index+1, sequence, result) 
}