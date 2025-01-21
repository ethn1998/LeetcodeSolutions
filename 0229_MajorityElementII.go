func majorityElement(nums []int) []int {
    counts := make(map[int]float64)
    fout := make([]int,0)
    for _,n := range nums{
        counts[n] += 1.0
    }
    Omega := float64(len(nums))/3.0
    for key, val := range(counts){
        if val > Omega{
            fout = append(fout,key)
        }
    }
    return fout
}
