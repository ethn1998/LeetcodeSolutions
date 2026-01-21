func minBitwiseArray(nums []int) []int {
    gen := []int{2,3,5,7,13,17,19,31}
    mem := make(map[int]bool)
    for _,p := range gen {
        mem[1 << p - 1] = true
    }
    ans := make([]int,len(nums))
    for i, n := range nums {
        if n == 2 {
            ans[i] = -1
        } else if mem[n] {
            ans[i] = n >> 1
        } else {
            //observation: answer always differs from input by power of 2.
            diff := 1
            for diff < n {
                k := n - diff 
                if k | (k + 1) == n { //to make more mathematically elegant
                    ans[i] = k
                }
                diff <<= 1
            }
        }
    }
    return ans
}
