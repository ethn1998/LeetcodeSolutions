func punishmentNumber(n int) int {
    /*
    This reminds me of my Math Olympiad days.
    Want: Sequence which satisfies integer partitioning condition.
    */
    var fout int
    for i := 1; i <= n; i++ {
        if existsPartition(i*i, i){ //Use another function to check if i can be partitioned
            fout += i*i
        }
    }
    return fout
}

func existsPartition(k int, target int) bool {
    if target < 0 || k < target {
        return false
    }
    if k == target {
        return true
    }
    //Extract digits; divide and conquer
    return existsPartition(k / 10, target - k % 10) || existsPartition(k / 100, target - k % 100) || existsPartition(k / 1000, target - k % 1000)
    
}
