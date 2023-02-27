func lengthOfLastWord(s string) int {
    count := 0
    check := false
    for i := len(s) - 1; i > -1; i-- {
        if s[i] == ' ' && check == true {
            return count
        }
        if s[i] != ' ' {
            count++;
            check = true
        }
    }
    return count
}