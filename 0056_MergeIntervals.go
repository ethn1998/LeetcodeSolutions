func merge(intervals [][]int) [][]int {
    //STEP 1: sort
    sort.Slice(intervals, func(i,j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    //STEP 2: merge intervals
    fout := make([][]int,0)
    fout = append(fout,intervals[0])
    for _, interval := range intervals {
        if fout[len(fout)-1][1] < interval[0] { //if current interval ends before candidate
            fout = append(fout,interval)
        } else {
            fout[len(fout)-1][1] = max(fout[len(fout)-1][1],interval[1])
        }
    }

    return fout
}
