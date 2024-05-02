func searchMatrix(matrix [][]int, target int) bool {
    rowStart, rowEnd := 0, len(matrix) - 1
    colStart, colEnd := 0, len(matrix[0]) - 1
    
    for rowStart <= rowEnd {
        rowMid := (rowStart + rowEnd) / 2
        
        if matrix[rowMid][0] == target {
            return true
        }
        
        if matrix[rowMid][0] < target {
            if matrix[rowMid][colEnd] >= target {
                for colStart <= colEnd {
                    colMid := (colStart + colEnd) / 2
                    
                    if matrix[rowMid][colMid] == target {
                        return true
                    }
                    
                    if matrix[rowMid][colMid] < target {
                        colStart = colMid + 1
                    } else {
                        colEnd = colMid - 1
                    }
                }
                break
            } else {
                rowStart = rowMid + 1
            }
        } else {
            rowEnd = rowMid - 1
        }
    }
    
    return false
}
