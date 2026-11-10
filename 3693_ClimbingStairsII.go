//greedy dp

func climbStairs(n int, costs []int) int {
    dp := make([]int,n+1) //minimum cost to reach step j
    for j := 1; j < n+1; j++ {
        dp[j] = dp[j-1] + costs[j-1] + 1
        if j > 1 {
            if dp[j] > dp[j-2] + costs[j-1] + 4 {
                dp[j] = dp[j-2] + costs[j-1] + 4
            }
        }
        if j > 2 {
            if dp[j] > dp[j-3] + costs[j-1] + 9 {
                dp[j] = dp[j-3] + costs[j-1] + 9
            }
        }
        //fmt.Println(dp)
    }
    return dp[n]
}
