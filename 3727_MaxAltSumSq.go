/*
Suppose we have n elements
Case 1: n = 2k
Then there are k +ve contributions and k -ve contributions
Case 2: n = 2k + 1
Then there are (k + 1) +ve contributions and k -ve contributions
First k elements smallest in magnitude
*/

func maxAlternatingSum(nums []int) int64 {
    n := len(nums)
    m := n / 2
    nums2 := make([]int64,n)
    for i,v := range nums {
        nums2[i] = int64(v*v)
    }
    slices.Sort(nums2)
    var ans int64
    for i,v := range nums2 {
        if i < m {
            ans -= v
        } else {
            ans += v
        }
    }
    return ans
}
