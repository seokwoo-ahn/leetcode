func minPathCost(grid [][]int, moveCost [][]int) int {
    row := len(grid)
    column := len(grid[0])
    
    cost := make([][]int, row)
    for i := 0; i < row; i++ {
        cost[i] = make([]int, column)
    }
    
    for i := 0; i < row; i++ {
        for j:= 0; j < column; j++ {
            cost[i][j] = math.MaxInt
        }
    }
    
    for i := 0; i < column; i++ {
        valNode := grid[0][i]
        cost[0][i] = valNode
    }
    
    for i := 0; i < row - 1 ; i++ {
        for j := 0; j < column; j++ {
            for k :=0; k < column; k++ {
                if cost[i+1][k] > grid[i+1][k] + cost[i][j] + moveCost[grid[i][j]][k] {
                    cost[i+1][k] = grid[i+1][k] + cost[i][j] + moveCost[grid[i][j]][k]
                }   
            }
        }
    }
    
    res := math.MaxInt
    for i := 0; i < column; i++ {
        res = int(math.Min(float64(res), float64(cost[row-1][i])))
    }
    return res
}
