func minPathSum(grid [][]int) int {
    column := len(grid[0])
    row := len(grid)
    cost := make([][]int, row)
    for i:= 0; i< row; i++ {
        cost[i] = make([]int, column)
    }
    
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++{
            cost[i][j] = math.MaxInt;
        }
    }
    
    cost[0][0] = grid[0][0];
    
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if j < len(grid[0]) - 1 {                
                if cost[i][j + 1] > cost[i][j] + grid[i][j + 1] {
                    cost[i][j + 1] = cost[i][j] + grid[i][j + 1];
                }
            }
        
            if i < len(grid) - 1 {
                if cost[i + 1][j] > cost[i][j] + grid[i + 1][j] {
                    cost[i + 1][j] = cost[i][j] + grid[i + 1][j];
                }
            }  
        }
    }
    return cost[len(grid) - 1][len(grid[0]) - 1];
}
