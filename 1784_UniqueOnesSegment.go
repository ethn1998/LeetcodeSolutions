func checkOnesSegment(s string) bool {
    var b0,b1 bool
    for _,char := range s {
        if b0 {
            if char == '0' {
                b1 = true
            } else if char == '1' && b1 {
                return false
            }
        } else if char == '1' {
            b0 = true
        }
    }
    return true
}
