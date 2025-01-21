import(
    "fmt"
    "strings"
)

func frequencySort(s string) string {
    var maxFreq int
    var fout,maxFreqword string
    var maxFreqchar rune
    counts := make(map[rune]int)
    for _,char := range(s){
        counts[char]++
    }
    for len(counts) > 0 {
        maxFreq = 0
        for k,v := range(counts){
            if v > maxFreq{
                maxFreqchar = k
                maxFreq = v
            }
        }
        maxFreqword = strings.Repeat(string(maxFreqchar),maxFreq)
        fout = fmt.Sprintf("%s%s",fout,maxFreqword)
        delete(counts,maxFreqchar)
    }
    return fout
}
