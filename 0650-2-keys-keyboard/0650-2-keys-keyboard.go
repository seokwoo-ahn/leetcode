func minSteps(n int) int {
    res, prime := 0, 2
    for n > 1 {
        for n % prime == 0 {
            res += prime
            n = n / prime
        }
        prime++
    }
    return res
}