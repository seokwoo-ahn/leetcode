func closedIsland(grid [][]int) int {
    res := 0
    check := false
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 0 {
                dfs(&grid, i, j, &check)
                if !check {
                    res++
                }
                check = false
            }
        }
    }
    return res
}

func dfs(grid *[][]int, xIdx int, yIdx int, check *bool) {
    (*grid)[xIdx][yIdx] = 2
    if xIdx == len(*grid) - 1 || xIdx == 0 || yIdx == len((*grid)[0]) - 1 || yIdx == 0 {
        *check = true
    }
    
    if xIdx != 0 && (*grid)[xIdx - 1][yIdx] == 0 {
        dfs(grid, xIdx-1, yIdx, check)
    }
    
    if xIdx != len((*grid)) -1 && (*grid)[xIdx + 1][yIdx]  == 0 {
        dfs(grid, xIdx+1, yIdx, check)
    }
    
    if yIdx != 0 && (*grid)[xIdx][yIdx - 1]  == 0 {
        dfs(grid, xIdx, yIdx-1, check)
    }
    
    if yIdx != len((*grid)[0]) -1 && (*grid)[xIdx][yIdx + 1]  == 0 {
        dfs(grid, xIdx, yIdx+1, check)
    }
}
