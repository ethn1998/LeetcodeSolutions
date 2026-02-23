func hasAllCodes(s string, k int) bool {
    n := len(s)
    //number of permutations
    m := twoPow(k)
    mem := make([]bool,m)
    x := 0
    for i := 0; i < n; i++ {
        x = x << 1
        if x >= m {
            x = x % m
        }
        if s[i] == '1' {
            x++
        }
        if i >= k - 1 {
            mem[x] = true
        }
    }
    for _,b := range mem {
        if !b {
            return false
        }
    }
    return true
}

func twoPow(n int) int {
    if n == 0 {
        return 1
    } else if n == 1 {
        return 2
    }
    x := twoPow(n/2)
    if n % 2 == 0 {
        return x * x
    }
    return 2 * x * x
}
