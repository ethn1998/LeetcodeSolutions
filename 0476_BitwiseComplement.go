func findComplement(num int) int {
    var fout int
    p := 1 //2^p
    for num > 0 {
        fout += p * ((num & 1)^1)
        p <<= 1
        num >>=1
    }
    return fout
}
