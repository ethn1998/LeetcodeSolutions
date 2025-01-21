func selfDividingNumbers(left int, right int) []int {
    var tenp,d int
    var isSelfDividing bool
    fout := make([]int,0)
    for n:=left; n<=right; n++{
        isSelfDividing = true
        tenp = 1
        for tenp < n{
            d = (n/tenp)%10
            if d != 0 { //Sanity check
                if n%d != 0{
                    isSelfDividing = false
                    break
                }
            } else {
                isSelfDividing = false
                break
            }
            tenp *= 10
        }
        if isSelfDividing{
            fout = append(fout,n)
        }
    }
    return fout
}
