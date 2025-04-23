func numRabbits(answers []int) int {
    var fout int
    freq := make(map[int]int)
    for _, ans := range(answers){
        freq[ans]++
    }
    for key, val := range(freq){
        for val > 0 {
            fout += key + 1
            val -= key + 1
        }
    }
    return fout
}
