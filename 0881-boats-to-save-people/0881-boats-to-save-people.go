func numRescueBoats(people []int, limit int) int {
    res := 0
    sort.Ints(people)
    
    firstIdx := 0
    lastIdx := len(people) - 1
    
    for firstIdx <= lastIdx {
        if firstIdx == lastIdx {
            return res + 1
        }
        
        if people[firstIdx] + people[lastIdx] <= limit {
            firstIdx++
            lastIdx--
        } else {
            lastIdx--
        }
        res++
    }
    return res
}
