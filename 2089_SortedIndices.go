func targetIndices(nums []int, target int) []int {
    var l,r int //count entries left and right of target
    for _,n := range nums {
        if n < target {
            l++
        } else if n > target {
            r++
        }
    }
    n := len(nums)
    fout := make([]int,0)
    for i := l; i < n - r; i++ {
        fout = append(fout,i)
    }
    return fout

}
