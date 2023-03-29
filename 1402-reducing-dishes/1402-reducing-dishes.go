func maxSatisfaction(satisfaction []int) int {
    sort.Ints(satisfaction)
    sum := 0
    idx := 0
    for i := len(satisfaction) - 1; i >= 0; i-- {
        sum += satisfaction[i]
        if sum < 0 {
            idx = i + 1
            break
        }
    }
    
    resArr := satisfaction[idx:]
    res := 0
    for i := 0; i < len(resArr); i++ {
        res += resArr[i]*(i+1)
    }
    return res
}