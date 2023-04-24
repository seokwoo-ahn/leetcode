func lastStoneWeight(stones []int) int {
    for len(stones) > 1 {
        sort.Ints(stones)
        if stones[len(stones) - 1] != stones[len(stones) - 2] {
            tmp := []int{stones[len(stones) - 1] - stones[len(stones) - 2]}
            stones = append(tmp, stones...)
        }
        stones = stones[:len(stones) - 2]
    }
    
    if len(stones) == 1{
        return stones[0]   
    } else {
        return 0
    }
}