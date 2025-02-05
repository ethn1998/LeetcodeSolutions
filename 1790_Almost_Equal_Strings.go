func areAlmostEqual(s1 string, s2 string) bool {
    /*
    The condition is satisfied if:
    1) s1 and s2 have the same letters
    2) s1 and s2 differ in at most two spots
    */
    if len(s1) != len(s2){ //This is redundant because of constraint len(s1) = len(s2)
        return false
    }
    //Store counts in list. This is better in this context because constraints -> only lowercase letters.
    var nmut int //Count number of differing indices
    var ascii1, ascii2 int //ascii integer values of runes
    s1counts := make([]int,26)
    s2counts := make([]int,26)
    for i := 0; i < len(s1); i++ {
        if s1[i] != s2[i] {
            nmut++
            if nmut > 2 {
                return false
            }
        }
        ascii1 = int(s1[i])
        ascii2 = int(s2[i])
        s1counts[ascii1-97]++
        s2counts[ascii2-97]++
    }
    for i,v := range(s1counts){ //Check if same letters
        if v != s2counts[i]{
            return false
        }
    }
    return true
}
