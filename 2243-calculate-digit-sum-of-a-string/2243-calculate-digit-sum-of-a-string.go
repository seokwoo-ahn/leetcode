func digitSum(s string, k int) string {
    for len(s) > k {
        sum := 0
        try := 0
        str := ""
        for i := 0; i < len(s); i++ {
            v, _ := strconv.Atoi(string(s[i]))
            sum = sum + v
            try++
            if try == k || i == len(s)-1 {
                str += strconv.Itoa(sum)
                sum = 0
                try = 0
            }
        }
         s = str
    }
    return s
}
