func makeConnected(n int, connections [][]int) int {
    visited := make([]bool, n)
    connMap := make(map[int][]int)

    if n - 1 > len(connections) {
        return -1
    }
    
    for i := 0; i < len(connections); i++ {
        conn := connections[i]
        connMap[conn[0]] = append(connMap[conn[0]], conn[1])
        connMap[conn[1]] = append(connMap[conn[1]], conn[0])
    }
    
    try := 0
    for i := 0; i < len(visited); i++ {
        if !visited[i]{
            try++
            dfs(i, connMap, &visited)
        }
    }
    return try - 1
}

func dfs(curNode int, conn map[int][]int, visited *[]bool) {
    (*visited)[curNode] = true
    
    if conn[curNode] != nil {
        for _, v := range conn[curNode]{
            if !(*visited)[v] {
                dfs(v, conn, visited)
            }
        }   
    }
}