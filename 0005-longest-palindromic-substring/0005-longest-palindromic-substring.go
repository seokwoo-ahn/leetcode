func longestPalindrome(s string) string {
    length := 1
    result := string(s[0])
    
    for i := 0; i < len(s); i++ {
        if length > len(s) - i {
            break;
        }
        temp := string(s[i])
        for j := i+1; j < len(s); j++ {
            temp = temp + string(s[j])
            if isPalindrome(temp){
                if len(temp) > length {
                    length = len(temp)
                    result = temp
                }
            }
        } 
    }
    return result
}

func isPalindrome(s string) bool {
    start := 0
    end := len(s) - 1
    for start < end {
        if s[start] != s[end]{
            return false   
        }
        start++
        end--
    }
    return true
}