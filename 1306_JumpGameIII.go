//bfs

func canReach(arr []int, start int) bool {
    n := len(arr)
    memo := make([]bool,n)
    q := []int{start}
    var src, dst int
    for i := 0; i < len(q); i++ {
        src = q[i]
        if arr[src] == 0 {
            return true
        } else {
            dst = src + arr[src]
            if dst < n {
                if !memo[dst] {
                    q = append(q,dst)
                    memo[dst] = true
                }
            }
            dst = src - arr[src]
            if dst >= 0 {
                if !memo[dst] {
                    q = append(q,dst)
                    memo[dst] = true
                }
            }
        }
    }
    return false
}
