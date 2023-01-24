type Move struct{
    Position int
    Moves int
}

func snakesAndLadders(board [][]int) int {
    length := len(board)
    max := length*length
    
    queue := []Move{{1,0}}
    
    visited := make([]bool, max)
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        position := current.Position
        if visited[position] {
            continue
        }
        
        visited[position] = true
        
        maxRange := int(math.Min(6, float64(max - position)))
        
        for i := 1; i <= maxRange; i++ {
            tempPosition := position + i - 1
            row := tempPosition/length
            column := tempPosition%length
            var value int
            
            if row % 2 == 0 {
                value = board[length-1-row][column]                
            } else {
                value = board[length-1-row][length-1-column]                
            }
            
            var newMove Move
            
            if value != -1 {
                newMove.Position = value
            } else {
                newMove.Position = tempPosition+1
            }
            
            if newMove.Position == max {
                return current.Moves + 1
            }
            
            if visited[newMove.Position] {
                continue
            }
            
            newMove.Moves = current.Moves + 1
            
            queue = append(queue, newMove)  
        }
    }
    return -1
}