func lengthOfLongestSubstring(s string) int {
    var result int
    var start int
    var temp string
    
    for i := 0; i < len(s); i++ {
        if strings.Contains(temp, string(s[i])) {
            index := strings.IndexByte(temp, s[i])
            start = start + index + 1
            temp = temp[index+1:]
        }
        result = int(math.Max(float64(result), float64(i - start + 1)))
        temp = temp + string(s[i])
    }
    return result
}
