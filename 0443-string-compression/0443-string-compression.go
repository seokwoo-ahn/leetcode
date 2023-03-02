func compress(chars []byte) int {
    if len(chars) == 1 {
        return 1
    }
    
    temp := chars[0]
    tempIdx := 0
    writeIdx := 0
    
    for i := 1; i < len(chars); i++ {
        if chars[i] != temp {
            chars[writeIdx] = temp
            writeIdx++
            count := i - tempIdx
            if count != 1 {
                countStr := strconv.Itoa(count)
                for j := 0; j < len(countStr); j++ {
                    chars[writeIdx] = countStr[j]
                    writeIdx++
                }
            }
            tempIdx = i
            temp = chars[i]
        }
    }
    
    count := len(chars) - tempIdx
    chars[writeIdx] = temp
    writeIdx++
    
    if count != 1 {
        countStr := strconv.Itoa(count)
        for j := 0; j < len(countStr); j++ {
            chars[writeIdx] = countStr[j]
            writeIdx++
        }
    }
    
    return writeIdx
}
