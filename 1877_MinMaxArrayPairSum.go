func minPairSum(nums []int) int {
    slices.Sort(nums)
    // fmt.Println(nums) //for debugging
    i := 0
    j := len(nums) - 1
    var fout int // initialize following problem constraints
    for i < j {
        if fout < nums[i] + nums[j] {
            fout = nums[i] + nums[j]
        }
        i++
        j--
    }
    return fout
}
