func maxSumOfSquares(num int, sum int) string {
    if sum > 9 * num {
        return ""
    }
    q := sum / 9 //number of leading nines
    r := sum % 9 //remainder. This will be strictly between 1 and 8 inclusive.
    var s int
    var mid string
    for i := 1; i <= r; i++ {
        if i * (r - i) >= s {
            mid = fmt.Sprintf("%d%d",i,r-i)
        }
    }
    pre := strings.Repeat("9",q) + strings.TrimRight(mid,"0")
    z := num - len(pre) //number of trailing zeroes needed
    return(pre + strings.Repeat("0",z))
}
