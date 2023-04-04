func partitionString(s string) int {
    res := 1
    m := make(map[byte]bool)
    idx := 0
    
    for idx < len(s) {
        b := s[idx]
        if m[b] {
            m = make(map[byte]bool)
            res++
        }
        m[b] = true
        idx++
    }

    return res
}
