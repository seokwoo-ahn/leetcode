func minScore(n int, roads [][]int) int {
    check := make([]bool, n)
    check[0] = true

    res := math.MaxInt
    
    for roads != nil {
        change := false
        for i := 0;  i < len(roads); i++ {
            edge := roads[0]
            roads = roads[1:]
            if check[edge[0] - 1] || check[edge[1] - 1] {
                res = int(math.Min(float64(edge[2]), float64(res)))
                check[edge[0] -1] = true
                check[edge[1] -1] = true
                change = true
            } else {
                roads = append(roads, edge)
            }   
        }
        if !change {
            return res
        }
    }
    return res
}
