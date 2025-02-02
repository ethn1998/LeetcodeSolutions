func check(nums []int) bool {
    /*
    IDEA: 
    Greedy algorithm to put minimum value at start of array, then check if resulting array is sorted.
    */
    if len(nums) <= 1 { //Degenerate case
        return true
    }
    minval := 101 //Initialize minimum value seen using constraints.
    for i := 0; i < len(nums); i++ {
        if nums[i] <= minval {
            minval = nums[i]
            v := nums[i]
            for j := 1; j < len(nums); j++ {
                k := (i+j) % len(nums) //Rotate index in original array
                if nums[k] < v { //If decrease detected
                    break
                } else { //If non-decreasing
                    if j < len(nums) - 1 { //If 
                        v = nums[k] //Update reference value
                    } else {
                        return true //If final index
                    }
                }
            }
        }
    }
    return false
}
