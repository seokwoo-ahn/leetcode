func validateStackSequences(pushed []int, popped []int) bool {
    q := make([]int, 0)
    idx := 0
    for i:= 0; i < len(pushed); i++ {
        q = append(q, pushed[i])
        for len(q) > 0 {
            if popped[idx] == q[len(q)-1] {
                q = q[:len(q)-1]
                idx++
            } else {
                break
            }
        }
    }
    
    return len(q) == 0
}