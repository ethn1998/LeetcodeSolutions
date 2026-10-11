func sumOfSquares(nums []int) int {
    var fout int
    n := len(nums)
    for i, m := range nums {
        if n % (i+1) == 0 {
            fout += m*m
        }
    }
    return fout
}
