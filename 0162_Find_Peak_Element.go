func findPeakElement(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
    for i,v := range(nums){
        if i > 0 && i < len(nums)-1 {
            if v > nums[i-1] && v > nums[i+1]{
                return i
            }
        } else if i == 0 {
            if nums[0] > nums[1]{
                return 0
            }
        } else {
            if nums[len(nums)-1] > nums[len(nums)-2]{
                return len(nums)-1
            }
        }
    }
    return 0
}
