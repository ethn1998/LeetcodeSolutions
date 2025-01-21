func numJewelsInStones(jewels string, stones string) int {
    var fout int
    jmap := make(map[rune]bool) //Go doesn't have hashsets, so map[type]bool would have to do
    for _,j := range(jewels){
        jmap[j] = true //Value is actually redundant here because we can use ok method to check key existence
    }
    for _,s := range(stones){
        _,ok := jmap[s]
        if ok {
            fout++
        }
    }
    return fout
}
