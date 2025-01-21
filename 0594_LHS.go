func findLHS(nums []int) int {
    var fout int
    counts := make(map[int]int)
    for _,n := range nums{
        counts[n]++
    }
    for k,v := range counts{
        u,ok := counts[k+1]
        if ok {
            if u+v > fout{
                fout = u+v
            }
        }
    }
    return fout
}
