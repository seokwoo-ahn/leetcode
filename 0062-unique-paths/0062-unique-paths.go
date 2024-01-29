import "math/big"

func uniquePaths(m int, n int) int {
    res := big.NewInt(1)
    if m > n {
        for i := m+n-2; i > m - 1; i-- {
            res = new(big.Int).Mul(res, big.NewInt(int64(i)))
        }
        for i := n - 1; i > 1; i-- {
            res = new(big.Int).Div(res, big.NewInt(int64(i)))
        }
    } else {
        for i := m+n-2; i > n - 1; i-- {
            res = new(big.Int).Mul(res, big.NewInt(int64(i)))
        }
        for i := m - 1; i > 1; i-- {
            res = new(big.Int).Div(res, big.NewInt(int64(i)))
        }
    }
    return int(res.Int64())
}