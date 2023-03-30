func areAlmostEqual(s1 string, s2 string) bool {
    count := 0
    idxArr := make([]int, 0)
    
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            if count == 2{
                return false
            }
            idxArr = append(idxArr, i)
            count++
        }
    }

    if count == 0 {
        return true
    } else if count == 1{
        return false
    } else if s1[idxArr[0]] == s2[idxArr[1]] && s2[idxArr[0]] == s1[idxArr[1]] {
            return true
    } else {
        return false
    }
}