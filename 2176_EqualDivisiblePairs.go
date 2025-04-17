func countPairs(nums []int, k int) int {
    //Try Brute Force
    var fout int
    seen := make(map[int]bool)
    for i, v := range(nums) {
        if !seen[v] {
            for j := i + 1; j < len(nums); j++ {
                if nums[j] == v {
                    if (i*j) % k == 0 {
                        fout++
                    }
                }
            }
        }
    }
    return fout
}
