func countOdds(low int, high int) int {
    if low & 1 == 0 {
        low++
    }
    
    if low > high {
        return 0
    }
    
    return (high - low)/2 + 1
}