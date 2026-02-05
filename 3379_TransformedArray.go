func constructTransformedArray(nums []int) []int {
    var pos int
    n := len(nums)
    result := make([]int,n)
    copy(result,nums)
    for i, val := range nums {
        pos = mod(i + val, n)
        result[i] = nums[pos]
    }
    return result
}

func mod(p,q int) int { //coerce into [0,q)
    for p < 0 {
        p += q
    }
    return p % q
}
