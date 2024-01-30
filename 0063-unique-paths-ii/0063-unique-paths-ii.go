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
        if obstacleGrid[i][0] == 0 {
            res[i][0] = 1
        } else {
            break
        }
    }
    
    for i := 0; i < height; i++ {
        if obstacleGrid[0][i] == 0{
            res[0][i] = 1
        } else {
            break
        }
    }

    for i := 1; i < length; i++ {
        for j := 1; j < height; j++ {
            if obstacleGrid[i][j] == 0 {
                temp := 0
                if obstacleGrid[i-1][j] == 0 {
                    temp += res[i-1][j] 
                }
                if obstacleGrid[i][j-1] == 0 {
                    temp += res[i][j-1]
                }
                res[i][j] = temp
            } else {
                res[i][j] = 0
            }
        }
    }
    return res[length-1][height-1]
}