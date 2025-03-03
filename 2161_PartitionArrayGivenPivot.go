func pivotArray(nums []int, pivot int) []int {
    left := make([]int,0)
    mid := make([]int,0)
    right := make([]int,0)
    for _,num := range(nums){
        if num < pivot {
            left = append(left,num)
        } else if num == pivot {
            mid = append(mid,num)
        } else {
            right = append(right,num)
        }
    }
    fout := make([]int,0)
    fout = append(fout,left...)
    fout = append(fout,mid...)
    fout = append(fout,right...)
    return fout
}
