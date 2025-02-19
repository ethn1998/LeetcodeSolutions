func strWithout3a3b(a int, b int) string {
    //wlog assume a > b
    var i int
    var fout string
    if a > b {
        for a > b {
            if i != 2 {
                fout += "a"
                a--
                i++
            } else{
                fout += "b"
                b--
                i = 0
            }
        }
    } else if a < b {
        for b > a {
            if i != 2 {
                fout += "b"
                b--
                i++
            } else {
                fout += "a"
                a--
                i = 0
            }
        }
    }
    // a = b case
    for a > 0 {
        if len(fout) > 0 {
            char := string(fout[len(fout)-1]) //Get final element of output
            if char == "a"{
                fout += "ba"
            } else {
                fout += "ab"
            }
            a--
        } else { //Empty string
            fout = "ab"
            a--
            b--
        }
    }
    return fout
}
