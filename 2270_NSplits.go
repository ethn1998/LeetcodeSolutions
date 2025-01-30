func waysToSplitArray(nums []int) int {
    var runsum0, fout int 
    runsum := make([]int,len(nums))
    for i,v := range(nums) {
        runsum0 += v
        runsum[i] = runsum0 
    }

    for i, s := range(runsum){
        if s >= runsum0 - s && i < len(nums)-1 {
            fout += 1
        }
    }
    return fout
}
