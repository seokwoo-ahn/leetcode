func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
    for i, v := range nums {
        if index, exist := m[target - v]; exist {
            return []int{i, index}
        }
        m[v] = i
    }
    return []int{}
}