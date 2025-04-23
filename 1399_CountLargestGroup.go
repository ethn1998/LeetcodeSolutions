func countLargestGroup(n int) int {
    groupsize := make(map[int]int)
    for i := 1; i <= n; i++ {
        m := i
        digitsum := 0
        for m > 0 {
            digitsum += m % 10
            m /= 10
        }
        groupsize[digitsum]++
    }
    var fout, maxsize int
    for _,val := range(groupsize){
        if val > maxsize {
            fout = 1
            maxsize = val
        } else if val == maxsize {
            fout++
        }
    }
    return fout
}
