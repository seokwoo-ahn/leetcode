func partitionString(s string) int {
    res := 1
    m := make([]bool, 26)
    idx := 0
    
    for idx < len(s) {
        b := s[idx]
        if m[b - 'a'] {
            m = make([]bool, 26)
            res++
        }
        m[b - 'a'] = true
        idx++
    }

    return res
}
