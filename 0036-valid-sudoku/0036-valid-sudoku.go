func isValidSudoku(board [][]byte) bool {
    checkColumn := make([](map[byte]bool),len(board))
    checkRow := make([](map[byte]bool), len(board))
    checkSection := make([](map[byte]bool), (len(board)/3)*(len(board)/3))
    
    for i := 0; i < len(board); i++ {
        checkColumn[i] = make(map[byte]bool)
        checkRow[i] = make(map[byte]bool)
    }
    
    for i := 0; i < (len(board)/3)*(len(board)/3); i++ {
        checkSection[i] = make(map[byte]bool)
    }
    
    for i := 0; i < len(board); i++ {
        for j :=0; j < len(board); j++ {
            value := board[i][j]
            if value == '.'{
                continue
            }
            sectionNum := i/3 + 3*(j/3)
            if checkRow[i][value] || checkColumn[j][value] || checkSection[sectionNum][value] {
                return false
            }
            
            checkRow[i][value] = true
            checkColumn[j][value] = true
            checkSection[sectionNum][value] = true
        }
    }
    return true
}
