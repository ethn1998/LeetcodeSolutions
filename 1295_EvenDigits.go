func findNumbers(nums []int) int {
    var fout,v int
    for _,n := range nums{
        v = int(math.Log10(float64(n)))
        if v%2 == 1{
            fout += 1
        }
    }
    return fout
}
