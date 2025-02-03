func maxAscendingSum(nums []int) int {
    var d int //Difference between consecutive elements
    maxAscSum := nums[0] //Record highest ascending sum seen
    ascSum := nums[0] //Initialize current ascending sum
    for i := 0; i < len(nums)-1; i++{
        d = nums[i+1] - nums[i]
        if d > 0 {
            ascSum += nums[i+1] 
        } else {
            ascSum = nums[i+1]
        }
        if ascSum > maxAscSum {
            maxAscSum = ascSum
        } 
    }
    return maxAscSum
}
