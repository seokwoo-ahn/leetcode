func divide(dividend int, divisor int) int {
    sign := (dividend < 0) == (divisor < 0)
    
    if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

    res := 0
    absDividend := int(math.Abs(float64(dividend)))
    absDivisor :=  int(math.Abs(float64(divisor)))
	for absDividend >= absDivisor {
		tempSub := absDivisor
        tempRes := 1
		for absDividend >= tempSub<<1 {
			tempSub <<= 1
			tempRes <<= 1
		}
		absDividend -= tempSub
		res += tempRes
	}

	if !sign {
		return -res
    } else {
        return res
    }
}