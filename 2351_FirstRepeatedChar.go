func repeatedCharacter(s string) byte {
    var fout byte
    seen := make(map[rune]bool)
    for _,char := range(s){
        if seen[char]{
            fout = byte(char)
            return fout
        }else{
            seen[char] = true
        }
    }
    return fout //Redundancy
}
