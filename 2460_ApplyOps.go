func applyOperations(nums []int) []int {
    n := len(nums) //Conserve length
    for i := 0 ; i < n-1; i++ {
        if nums[i] == nums[i+1] {
            nums[i] *= 2
            nums[i+1] = 0
        } //else do nothing
    }
    for j := 0 ; j < len(nums); {
        if nums[j] == 0 {
            nums = append(nums[:j],nums[j+1:]...)
        } else {
            j++
        }
    }
    for len(nums) < n {
        nums = append(nums,0)
    }
    return nums
}
