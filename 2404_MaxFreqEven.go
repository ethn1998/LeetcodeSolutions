func mostFrequentEven(nums []int) int {
    var maxFreq int
    counts := make(map[int]int)
    for _,n := range nums{
        counts[n]++
    }
    fout := -1 //Default to -1
    for key,count := range counts {
        if key % 2 == 0 { //Only look at evens
            if count > maxFreq {
                maxFreq = count
                fout = key
            } else if count == maxFreq {
                if key < fout {
                    fout = key
                }
            }
        }
    }
    return fout 
}
