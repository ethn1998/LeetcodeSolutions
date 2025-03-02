func timeRequiredToBuy(tickets []int, k int) int {
    //Simulate process
    var i,t int
    n := len(tickets)
    for true { //While true
        if tickets[i] > 0 {
            tickets[i]--
            if i == k && tickets[k] == 0 {
                break
            }
            t++
        } //else do nothing
        i++
        i %= n
    }
    return t + 1
}
