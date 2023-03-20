func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {
        return true
    }
    count := 0
    for i := 0; i < len(flowerbed); i++ {
        left := i == 0 || flowerbed[i-1] == 0
        right := i == len(flowerbed) -1 || flowerbed[i+1] == 0
        
        if left && right && flowerbed[i] == 0{
            flowerbed[i] = 1
            count++
        }
        
        if count >= n {
            return true
        }
    }
    return count >= n
}
