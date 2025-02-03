func longestMonotonicSubarray(nums []int) int {
    var d int //Difference between consecutive elements
    incr := 1
    desc := 1
    maxstreak := 1
    for i := 0; i < len(nums)-1; i++ {
        d = nums[i+1] - nums[i]
        if d > 0 {
            incr += 1
            desc = 1
        } 
        if d < 0 {
            desc += 1
            incr = 1
        } 
        if d == 0 { //No change
            incr = 1
            desc = 1
        }

        if incr > desc {
            if incr > maxstreak {
                maxstreak = incr
            }
        }
        if desc > incr {
            if desc > maxstreak {
                maxstreak = desc
            }
        }
    }
    return maxstreak
}
