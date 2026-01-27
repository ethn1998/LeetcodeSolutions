//two scans of sorted array: once to find min, second to get pairs
func minimumAbsDifference(arr []int) [][]int {
    slices.Sort(arr)
    mindiff := 1 << 63 - 1
    i0 := 0
    for i := 0; i < len(arr) - 1; i++ {
        if mindiff > arr[i+1] - arr[i] {
            mindiff = arr[i+1] - arr[i]
            i0 = i
        }
    }
    fout := make([][]int,0)
    for i := i0; i < len(arr) - 1; i++ {
        if arr[i+1] - arr[i] == mindiff {
            fout = append(fout, []int{arr[i], arr[i+1]})
        }
    }
    return fout
}
