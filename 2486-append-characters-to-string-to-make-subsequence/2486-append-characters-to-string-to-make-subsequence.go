func appendCharacters(s string, t string) int {
    for i := 0; i < len(s); i++ {
        if len(t) == 0 {
            return 0
        }
        if t[0] == s[i] {
            t = t[1:]
        }
    }
    
    return len(t)
}