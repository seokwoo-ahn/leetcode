func generateParenthesis(n int) []string {
    var result []string
    if n == 0 {
        return result
    }
    
    queue := make([]*Pattern, 0)
    
    init := &Pattern{
        Left: n,
        Right: n,
        Str: "",
    }
    
    queue = append(queue, init)
    
    for len(queue) >0 {
        q := queue[0]
        queue = queue[1:]
        
        if q.Left == 0 && q.Right == 0 {
            result = append(result, q.Str)
            continue
        } 
        
        if q.Left == 0 {
            q.Right = q.Right - 1
            q.Str = q.Str + ")"
            queue = append(queue, q)
        } else if q.Left < q.Right{
            right := &Pattern{
                Left: q.Left,
                Right: q.Right - 1,
                Str: q.Str + ")",
            }
            
            left := &Pattern{
                Left: q.Left - 1,
                Right: q.Right,
                Str: q.Str + "(",
            }
            queue = append(queue, right)
            queue = append(queue, left)
        } else {
            q.Left = q.Left - 1
            q.Str = q.Str + "("
            queue = append(queue, q)
        }
        
    }
    
    return result
}

type Pattern struct {
    Left int
    Right int
    Str string
}