func minimumDifference(nums []int, k int) int {
    if k == 1 { //common sense case
        return 0
    }
    slices.Sort(nums)
    mindiff := 1 << 63 - 1
    for i := 0; i <= len(nums) -k; i++ {
        diff := nums[i+k-1] - nums[i]
        if diff < mindiff {
            mindiff = diff
        }
    }
    return mindiff
}
