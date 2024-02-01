import "fmt"

func minimumTotal(triangle [][]int) int {
    height := len(triangle)
    triSum := make([][]int, height)
    for i:= 0; i < height; i++ {
        triSum[i] = make([]int, i+1)
        for j := 0; j < i+1; j++ {
            triSum[i][j] = 1 << 31
        }
    }
    
    res := 1 << 31
    for i := 0; i < height; i++ {
        for j := 0; j < i+1; j++ {
            if i == 0 {
                triSum[i][j] = triangle[i][j]
            } 
            if i == height -1 {
                if res > triSum[i][j] {
                    res = triSum[i][j]
                }
            } else {
                if triSum[i+1][j] > triSum[i][j] + triangle[i+1][j] {
                    triSum[i+1][j] = triSum[i][j] + triangle[i+1][j] 
                }
                triSum[i+1][j+1] = triSum[i][j] + triangle[i+1][j+1]
            }
        }
    }
    
    fmt.Println(triSum)
    
    return res
}