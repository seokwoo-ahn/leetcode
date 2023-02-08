type Jump struct {
    position int
    count int
}

func jump(nums []int) int {
    if len(nums) < 2 {
        return 0
    }
    init := Jump{}
    
    check := make([]bool, len(nums))
    var queue []Jump
    
    queue = append(queue, init)
    
    for len(queue) != 0 {
        p := queue[0]
        
        queue = queue[1:]
        
        jump := nums[p.position]
        if jump == 0 {
            continue
        }
        
        for i := 1; i < jump + 1; i++ {
            var newPos Jump
            newPos.position = p.position + i
            if check[newPos.position] == true {
                continue
            }
            newPos.count = p.count + 1
            if newPos.position == len(nums) - 1{
                return newPos.count
            }
            queue = append(queue, newPos)
            check[newPos.position] = true
        }
    }
    return 0
}