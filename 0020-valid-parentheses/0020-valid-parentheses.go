func isValid(s string) bool {
    stack := make([]byte,0)
    
    for i := 0; i < len(s); i++ {
        if s[i] == '{' || s[i] == '(' || s[i] == '['{
            stack = append(stack, s[i])
        } else if len(stack) == 0 {
            return false
        } else if  s[i] == '}' {
            if stack[len(stack) - 1] == '{' {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        } else if s[i] == ')' {
            if stack[len(stack) - 1] == '(' {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        } else if s[i] == ']'{
            if stack[len(stack) - 1] == '[' {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        }
    }
    return len(stack) == 0
}
