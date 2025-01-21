func isIsomorphic(s string, t string) bool {
    var schar, tchar byte
    stot := make(map[byte]byte)
    ttos := make(map[byte]byte)
    for i := 0; i < len(s); i++{
        schar = s[i]
        tchar = t[i]
        val, ok := stot[schar]
        if ok {
            if val != tchar{
                return false
            }
        }else{
            stot[schar] = tchar
        }
        //Non-injective relationships are not allowed
        val, ok = ttos[tchar]
        if ok {
            if val != schar{
                return false
            }
        }else{
            ttos[tchar] = schar
        }
    }
    return true
}
