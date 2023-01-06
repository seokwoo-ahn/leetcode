func maxIceCream(costs []int, coins int) int {
    var result int
    sort.Ints(costs)
    
    for _, v := range costs {
        if (coins - v < 0 ){
            return result 
        }
        coins = coins -v 
        result++
    }
    return result
}