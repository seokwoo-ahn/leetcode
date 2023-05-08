func diagonalSum(mat [][]int) int {
    sum := 0
    for i := 0; i < len(mat); i++ {
        sum += mat[i][i] + mat[i][len(mat) - i - 1]
    }
    
    if len(mat)%2 != 0 {
        sum -= mat[len(mat)/2][len(mat)/2]
    }
    return sum
}