func findMissingAndRepeatedValues(grid [][]int) []int {
    //This is probably inefficient but I think it works
    //Hashset for memorizing seen numbers
    var a,b,n2 int
    mem := make(map[int]bool)
    for _,row := range(grid) {
        for _,v := range(row) {
            _, ok := mem[v]
            if ok {
                a = v
            } else {
                mem[v] = true
            }
            n2++
        }
    }
    for i := 1; i <= n2; i++ {
        _, ok := mem[i]
        if !ok {
            b = i
            break
        }
    }
    return []int{a,b}
}
