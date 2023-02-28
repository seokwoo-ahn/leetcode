func mySqrt(x int) int {
    first := 0
    last := x
    
    for first <= last {
        mid := (first + last)/2
        if mid*mid > x {
            last = mid - 1
        } else if mid*mid < x {
            first = mid + 1
        } else {
            return mid
        }
    }
    
    return last
}