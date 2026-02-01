/*
This is first element of nums + sum of two smallest from remaining nums.
sort and get leading elements: O(nlogn)
try single sweep in O(n) time?
*/

func minimumCost(nums []int) int {
    //init following problem constraints
    var n int
    a := 50
    b := 50
    for i := 1; i < len(nums); i++ {
        n = nums[i]
        if n < a {
            b = a
            a = n
        } else if n < b {
            b = n
        }
    }
    return nums[0] + a + b
}
