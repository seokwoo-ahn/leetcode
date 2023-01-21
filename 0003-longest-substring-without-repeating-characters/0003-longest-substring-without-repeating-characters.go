func lengthOfLongestSubstring(s string) int {
    var result int
    for i := 0; i < len(s); i++ {
        if (len(s) - i) <= result{
            break
        }
        if result < findSubStringLen(s, i){
            result = findSubStringLen(s, i)
        }  
    }
    return result
}

func findSubStringLen(s string, start int) int {
    var subString string
    for i := start; i < len(s); i++ {
        if strings.Contains(subString, string(s[i])) {
            return len(subString)
        }
       subString = subString + string(s[i])
    }
    return len(subString)
}
