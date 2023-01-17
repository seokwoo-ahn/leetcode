func minFlipsMonoIncr(s string) int {
    var result float64
    var zeroNum float64
    var change float64
   
    zeroNum = float64(strings.Count(s, "0"))
    result = zeroNum
    change = zeroNum

    for i := 0; i < len(s); i++ {
        if s[i] == '0'{
            change--
            result = math.Min(result, change)
        } else {
            change++
        }
    }
    return int(result)
}