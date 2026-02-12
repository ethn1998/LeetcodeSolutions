func longestBalanced(s string) int {
    var fout, k int
    n := len(s)
    counts := make([]int,26)
    //i = 0
    for i,char := range s {
        k = int(char - 'a')
        counts[k]++
        if isBalanced(counts) {
            fout = i + 1
        } 
    }
    //i > 0
    for i := 1; i < n; i++ {
        k = int(s[i-1] - 'a')
        counts[k]--
        for j := n-1; j >= i + fout; j-- {
            k = int(s[j] - 'a')        
            counts[k]--
        }
        for j := i+fout ; j < n; j++ {
            k = int(s[j] - 'a')        
            counts[k]++
            if isBalanced(counts) {
                if j-i+1 > fout {
                    fout = j-i+1
                }
            }
        }
    }
    return fout
}

func isBalanced(counts []int) bool {
    var v0 int
    for _,v := range counts {
        if v > 0 {
            if v0 > 0 {
                if v != v0 {
                    return false
                }
            } else {
                v0 = v
            }
        }
    }
    return true
}
