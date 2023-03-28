func getRow(rowIndex int) []int {
    res := make([][]int, rowIndex+1)
    
    for i := 0; i <= rowIndex; i++ {
        res[i] = make([]int, i+1)
        for j := 0; j <= i; j++ {
            if j == 0 || j == i {
                res[i][j] = 1
            } else {
                res[i][j] = res[i-1][j-1] + res[i-1][j]
            }
        }
    }
    return res[rowIndex]
}
