func addBinary(a string, b string) string {
    carry := 0
    result := ""
    
    idxA := len(a) - 1
    idxB := len(b) - 1
    
    for idxA >= 0 || idxB >= 0 || carry != 0{
        var tempA byte
        var tempB byte
        sum := 0
        
        if idxA >= 0 {
            tempA = a[idxA]            
        } else {
            tempA = 0 
        }
        idxA--
        
        if idxB >= 0 {
            tempB = b[idxB]
        } else {
            tempB = 0
        }
        idxB--
        
        
        if tempA == '1' {
            sum++
        }
        
        if tempB == '1' {
            sum++
        }
        
        sum = carry + sum
        
        carry = sum/2
        result = strconv.Itoa(sum%2) + result
    } 
    return result
}