func findMaxK(nums []int) int {
    m := make(map[int]bool)
    res := -1
    
    for i := 0; i < len(nums); i++ {
        v := nums[i]
        
        if m[-v] == true {
            if v < 0 && -v > res {
                res = -v
            } else if v > res{
                res = v
            }
        } else {
            m[v] = true
        }  
    }
    
    return res
}