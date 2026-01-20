func minBitwiseArray(nums []int) []int {
    ans := make([]int,len(nums))
    for i, n := range nums {
        ans[i] = -1
        for k := 1; k < n; k++ { //brute force
            if k | (k + 1) == n {
                ans[i] = k
                break
            }
        }
    }
    return ans
}
