func maximumElementAfterDecrementingAndRearranging(arr []int) int {
    n := len(arr)
    slices.Sort(arr)
    for i := range arr {
        if i > 0 {
            if arr[i] > arr[i-1] + 1 {
                arr[i] = arr[i-1] + 1
            }
        }
        if arr[i] > i + 1 {
            arr[i] = i + 1
        }
    }
    return arr[n-1]
}
