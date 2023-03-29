func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    res := make([][]int, 0)
    dfs(candidates, 0, target, 0, make([]int,0), &res)
    return res
}

func dfs(candidates []int, idx int, target int, sum int, arr []int, res *[][]int) {
    for i := idx; i < len(candidates); i++ {
        if i != idx && candidates[i-1] == candidates[i]{
            continue
        }
        tempSum := sum + candidates[i]
        tempArr := make([]int, len(arr))
        copy(tempArr, arr)
        tempArr = append(tempArr, candidates[i])
        
        if tempSum < target {
            dfs(candidates, i+1, target, tempSum, tempArr, res)
        } else if tempSum == target {
            (*res) = append(*res, tempArr)
            return 
        } else {
            return 
        }
    }
}
