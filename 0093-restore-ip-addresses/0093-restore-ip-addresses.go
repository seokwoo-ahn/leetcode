import "strconv"

func restoreIpAddresses(s string) []string {
    var index int
    var count int
    var num int
    var valid string
    var value string
    var result []string
    
    findValidAddress(s, index, count, num, valid, value, &result)
    return result
}

func findValidAddress(s string, index int, count int, num int, valid string, value string, result *[]string){
    if index == len(s){
        if count < 4 && num == 3 {
            temp, _ := strconv.Atoi(value)
            if temp < 256{
                *result = append(*result, valid)
            }
        }
        return
    }
    
    if index == 0  {
        findValidAddress(s, index+1, count+1, num, string(s[0]), string(s[0]), result)
        return
    }
    
    if count < 3 && value[0] != '0'{
        appendString := valid + string(s[index])
        tempValue := value + string(s[index])
        findValidAddress(s, index+1, count+1, num, appendString, tempValue, result)
    }
    temp, _ := strconv.Atoi(value)
    if temp > 255{
        return
    }
    appendString := valid + "." + string(s[index])
    findValidAddress(s, index+1, 0, num+1, appendString, string(s[index]), result)
}