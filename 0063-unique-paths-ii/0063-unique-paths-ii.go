func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    length := len(obstacleGrid)
    height := len(obstacleGrid[0])
    
    if obstacleGrid[0][0] == 1 || obstacleGrid[length-1][height-1] == 1 {
        return 0
    }
    
    res := make([][]int, length)
    for i := 0; i < length; i++ {
        res[i] = make([]int, height)
        for j:= 0; j < height; j++ {
            res[i][j] = 0
        }
    }
    
    for i := 0; i < length; i++ {
        for j := 0; j < height; j++ {
            if obstacleGrid[i][j] == 0 {
                if i == 0 && j == 0 {
                    res[i][j] = 1
                }
                if i > 0 {
                    res[i][j] += res[i-1][j]
                }
                if j > 0 {
                    res[i][j] += res[i][j-1]
                }
            } else {
                res[i][j] = 0
            }
        }
    }
    return res[length-1][height-1]
}