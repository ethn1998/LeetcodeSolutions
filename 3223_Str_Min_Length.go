func minimumLength(s string) int {
    var fout int
    counts := make(map[rune]int)
    for _,char := range(s){
        counts[char]++
    }
    for _,val := range(counts){
        if val % 2 == 1{
            fout += 1
        } else {
            fout += 2
        }
    }
    return fout
}
