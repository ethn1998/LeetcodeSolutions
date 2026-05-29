func minElement(nums []int) int {
    fout := 36 //max output when input = 9999
    for _,n := range nums {
        var x int
        for n > 0 {
            x += n % 10
            n /= 10
        }
        if fout > x {
            fout = x
        }
    }
    return fout
}
