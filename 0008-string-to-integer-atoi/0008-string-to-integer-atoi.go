func myAtoi(s string) int {
    res := ""
    start := false
    sign := 1
    for i := 0; i < len(s); i++ {
        temp := rune(s[i])
        if !unicode.IsDigit(temp) && start == false && !check(temp) {
            return 0;
        }
        
        if !unicode.IsDigit(temp) && start == true {
            break;
        }
        
        if unicode.IsDigit(temp) {
            res = res + string(temp)
            start = true
        } else if temp == '-' || temp == '+'{
            if i == len(s) - 1 {
                break
            }
            if start == false && unicode.IsDigit(rune(s[i+1])) {
                if temp == '-' {
                    sign = -1
                }
            }
            start = true
        }
    }
    resInt, _ := strconv.Atoi(res)
    resInt = resInt*sign
    if resInt > math.MaxInt32{
        return math.MaxInt32
    } else if resInt < math.MinInt32{
         return math.MinInt32
    }
    return resInt
}

func check(input rune) bool {
    if input == ' ' || input == '-' || input == '+' {
        return true   
    }
    return false
}

