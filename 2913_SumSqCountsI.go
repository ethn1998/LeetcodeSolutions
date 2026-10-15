//O(n2). Inefficient.

func sumCounts(nums []int) int {
    var fout int
    n := len(nums)
    for l := 0; l < n; l++ {
        for r := l+1; r < n+1; r++ {
            mem := make(map[int]bool)
            for i := l; i < r; i++ {
                mem[nums[i]] = true
            }
            fout += len(mem) * len(mem)
        } 
    }
    return fout
}
