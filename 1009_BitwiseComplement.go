func bitwiseComplement(n int) int {
    if n == 0 { //DEGENERATE CASE
        return 1
    }
    var fout int
    p := 1 //2^p
    for n > 0 {
        fout += p * ((n & 1)^1)
        p <<= 1
        n >>=1
    }
    return fout
}
