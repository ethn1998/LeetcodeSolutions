func canReach(s string, minJump int, maxJump int) bool {
    n := len(s)
    if s[n-1] == '1' {
        return false
    }
    memo := make([]bool,n)
    memo[0] = true
    q := []int{0}
    for i := 0; i < len(q); i++ {
        src := q[i]
        l := src + minJump
        r := min(src + maxJump, n-1)
        for dst := r; dst >= l; dst-- { //greed
            if dst == n-1 {
                return true
            }
            char := s[dst]
            if char == '0' {
                if !memo[dst] {
                    q = append(q,dst)
                    memo[dst] = true
                }
            }
        }
    }

    return false
}
