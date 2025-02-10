func rev(n int) int { //Implement reverse number
    var fout int
    for n > 0 {
        r := n % 10
        fout *= 10
        fout += r
        n /= 10
    }
    return fout 
}

func countNicePairs(nums []int) int {
    // Apply strategy from count bad pairs
    counts := make(map[int]int)
    for _,v := range(nums) {
        counts[v-rev(v)]++
    }
    var fout int
    for _,val := range(counts) {
        fout += val*(val-1)
    }
    return fout/2 % 1000000007
}
