func arrayChange(nums []int, operations [][]int) []int {
    m := make(map[int]int)
    
    for i:= 0; i < len(nums); i++ {
        m[nums[i]] = i
    }
    
    for i:=0; i < len(operations); i++ {
        index, ok := m[operations[i][0]]
        if ok {
            nums[index] = operations[i][1]
            m[operations[i][1]] = index
        }
    }
    return nums
}
