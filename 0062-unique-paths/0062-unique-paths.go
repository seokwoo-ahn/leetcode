func uniquePaths(m int, n int) int {
    root := make([][]int, m)
    for i := 0; i < m; i++ {
        root[i] = make([]int, n)
        for j := 0; j < n; j++ {
            root[i][j] = 1
        }
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            root[i][j] = root[i-1][j] + root[i][j-1]
        }
    }
    return root[m-1][n-1]
}