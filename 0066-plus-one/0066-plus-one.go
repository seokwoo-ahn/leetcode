func plusOne(digits []int) []int {
    length := len(digits)
    
    for i := length -1; i > -1; i -- {
        if digits[i] < 9 {
            digits[i] = digits[i] + 1
            return digits
        } else {
            digits[i] = 0
        }
    }
    
    result := []int{1}
    result = append(result, digits...)
    return result
}