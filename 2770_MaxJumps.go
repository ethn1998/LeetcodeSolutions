func maximumJumps(nums []int, target int) int {
    n := len(nums)
    dp := make([]int,n) //count max number of jumps needed to reach index
    for i := 0; i < n-1; i++ {
        dp[i] = -1
    }
    for i := n-2; i >= 0; i-- {
        for j := i+1; j < n; j++ {
            if nums[j] - nums[i] <= target && nums[i] - nums[j] <= target {
                if dp[j] != -1 {
                    if dp[i] < dp[j] + 1 {
                        dp[i] = dp[j] + 1
                    }
                }
            }
        }
    }
    return dp[0]
}
