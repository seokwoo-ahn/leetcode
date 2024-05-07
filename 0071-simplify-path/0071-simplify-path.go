func simplifyPath(path string) string {
    stack := make([]string, 0)
    
    path = path + "/"

    dirName := ""
    for i := 0; i < len(path); i++ {
        if path[i] == '/' {
            if dirName == ".." {
                if len(stack) != 0 {
                    stack = stack[:len(stack)-1]
                }
            } else if dirName != "." && dirName != "" {
                stack = append(stack, dirName)
            }
            dirName = ""
        } else {
            dirName = dirName + string(path[i])
        }
    }
    
    res := ""
    if len(stack) == 0 {
        return "/"
    }
    for i := 0; i < len(stack); i++ {
        res = res + "/" + stack[i]
    }
    return res
}
