func diagonalSum(mat [][]int) int {
    sum := 0
    length := len(mat)
    for i := 0; i < length; i++ {
        sum += mat[i][i] + mat[i][length - i - 1]
    }
    
    if length%2 != 0 {
        sum -= mat[length/2][length/2]
    }
    return sum
}