func topKFrequent(nums []int, k int) []int {
    counts := make(map[int]int)
    for i := 0; i<len(nums); i++{
        counts[nums[i]]++
    }
    fout := make([]int,0,len(counts))
    for j := range(counts){
        fout = append(fout,j)
    }
    sort.SliceStable(fout, func(i, j int) bool{
        return counts[fout[i]] > counts[fout[j]]
    })
    return fout[0:k]
}
