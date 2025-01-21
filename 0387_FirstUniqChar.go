func firstUniqChar(s string) int {
    fout := 999999999
    seen := make(map[rune]int)
    mult := make(map[rune]bool)
    for i,char := range(s){
        _, ok := seen[char]
        if ok{
            mult[char] = true
        }else{
            seen[char] = i
        }
    }
    for key,val := range(seen){
        if !mult[key]{
            if val < fout{
                fout = val
            }
        }
    }
    if fout == 999999999 {
        return -1
    } else {
        return fout
    }
}
