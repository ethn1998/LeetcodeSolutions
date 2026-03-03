func findKthBit(n int, k int) byte {
    if n == 1 {
        return '0'
    }
    m := 1 << n
    fmt.Println("In:",k,m)
    if k < m / 2 {
        return findKthBit(n - 1, k)
    } else if k > m / 2 {
        return '0' + '1' - findKthBit(n - 1, m - k)
    }
    return '1'
}
