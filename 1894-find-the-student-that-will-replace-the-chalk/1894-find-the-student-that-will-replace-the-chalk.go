func chalkReplacer(chalk []int, k int) int {
    sum := 0
    for i := 0; i < len(chalk); i++ {
        sum += chalk[i]
    }
    
    k = k%sum
    
    for i := 0; i < len(chalk); i++ {
        k = k - chalk[i]
        if k < 0 {
            return i
        }
    }
    return -1
}
