func isTrionic(nums []int) bool {
    var decr, incr bool
    if nums[0] >= nums[1] {
        return false
    } 
    for i := 1; i < len(nums)-1; i++ {
        if nums[i] == nums[i+1] {
            return false
        }
        if !decr {
            if nums[i] > nums[i+1] {
                decr = true
            }
        } else {
            if nums[i] < nums[i+1]{
                incr = true
            }
        }
        if incr {
            if nums[i] > nums[i+1] {
                return false
            }
        }
    }
    if incr {
        return true
    }
    return false
}
