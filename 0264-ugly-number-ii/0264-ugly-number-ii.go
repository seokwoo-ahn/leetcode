func nthUglyNumber(n int) int {
    uArr := []int{1}
    idx2, idx3, idx5 := 0, 0, 0
    val2, val3, val5 := 2, 3, 5
    res := 1
    
    for i := 1; i < n; i ++ {
        res = min(val2, min(val3, val5))
        uArr = append(uArr, res)
        
        if res == val2 {
            idx2++
            val2 = 2 * uArr[idx2]
        }
        
        if res == val3 {
            idx3++
            val3 = 3 * uArr[idx3]
        }
        
        if res == val5 {
            idx5++
            val5 = 5 * uArr[idx5]
        }
    }
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}