func isArraySpecial(nums []int) bool {
    if len(nums) == 1 { //Degenerate case
        return true
    }
    var p int
    for i := 0; i < len(nums) - 1; i++ {
        p = nums[i] + nums[i+1]
        if p % 2 == 0 {
            return false
        }
    }
    return true
}
