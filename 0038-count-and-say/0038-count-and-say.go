func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    
    temp := countAndSay(n-1)
    fmt.Println(temp)
    res := ""
    tempValue := temp[0]
    count := 1
    for i := 1; i < len(temp); i++ {
        if tempValue != temp[i] {
            res += strconv.Itoa(count)
            res += string(tempValue)
            tempValue = temp[i]
            count = 1
        } else {
            count++
        }
    }
    res += strconv.Itoa(count)
    res += string(tempValue)
    
    return res
}
