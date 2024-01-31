func numDecodings(s string) int {
    if s[0] == '0'{
        return 0
    }

    length := len(s)

    dp := make([]int, length + 1)
    dp[length] = 1
    if s[length - 1] != '0'{
        dp[length - 1] = 1;
    }

    for i := length - 2; i >= 0; i--{
        if s[i] != '0'{
            dp[i] += dp[i + 1]
        }
        if s[i] == '1' || s[i] == '2' && s[i + 1] < '7'{
            dp[i] += dp[i + 2]
        } 
    }
    return dp[0]
}