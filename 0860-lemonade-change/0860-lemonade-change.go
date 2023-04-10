func lemonadeChange(bills []int) bool {
    numFiveD := 0
    numTenD := 0
    
    for i := 0; i < len(bills); i++ {
        if bills[i] == 5 {
            numFiveD++
        } else if bills[i] == 10 {
            if numFiveD > 0 {
                numFiveD--
                numTenD++
            } else {
                return false
            }
        } else {
            if numTenD > 0 && numFiveD > 0 {
                numTenD--
                numFiveD--
            } else {
                if numFiveD - 3 >= 0 {
                    numFiveD = numFiveD - 3
                } else {
                    return false
                }
            }
        }
    }
    return true
}
