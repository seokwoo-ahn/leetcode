func shipWithinDays(weights []int, days int) int {
    totalWeight := 0
    maxWeight := 0
    for i := 0; i < len(weights); i++ {
        totalWeight += weights[i]
        maxWeight = int(math.Max(float64(maxWeight), float64(weights[i])))
    }
    
    for {
        if isPossible(weights, days, maxWeight) {
            return maxWeight
        } else {
            maxWeight++
        }
    }
    return 0
}

func isPossible(weights []int, days int, capacity int) bool {
    dayNeed := 0
    temp := 0
    
    for i := 0; i < len(weights); i++ {
        if temp + weights[i] > capacity {
            temp = weights[i]
            dayNeed++
        } else {
            temp += weights[i]
        }
        
        if dayNeed >= days {
            return false
        }
    }
    return true
} 