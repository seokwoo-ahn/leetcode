func partition(s string) [][]string {
    var index int
    var substring string
    var palindrome []string
    var result [][]string
    
    findPalindrome(s, index, substring, palindrome, &result)
    return result
}

func findPalindrome(s string, index int, substring string, palindrome []string, result *[][]string) {
	if index == len(s) {
		if isPalindrome(substring) && substring != "" {
			tempPalindrome := make([]string, len(palindrome))
			copy(tempPalindrome, palindrome)
			tempPalindrome = append(tempPalindrome, substring)
			*result = append(*result, tempPalindrome)
		}
		return
	}

	candidate := substring + string(s[index])
	if isPalindrome(candidate) {
        tempPalindrome := append(palindrome, candidate)
		findPalindrome(s, index+1, "", tempPalindrome, result)
	}
	findPalindrome(s, index+1, candidate, palindrome, result)
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