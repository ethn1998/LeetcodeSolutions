func findTheDifference(s string, t string) byte {
    var fout byte
    var count int
    counts := make(map[rune]int)
    for _,char := range(s){
        counts[char]++
    }
    for _,char := range(t){
        count = counts[char]
        if count == 0{
            fout = byte(char)
            break //singular return must be outside loop
        }
        counts[char]--
    }
    return fout 
}
