func repeatedNTimes(nums []int) int {
    mem := make(map[int]int,len(nums)) //Without specifying capacity, the map will allocate new memory as it grows, contributing additional work to the garbage collector.
    //n := len(nums) / 2
    for _,k := range nums {
        mem[k]++
        //fmt.Println(k, mem[k])
        if mem[k] > 1 { //Exactly one element is duplicated n times
            return k
        }
    }
    return -1 //this can never happen, but golang functions demand this kind of redundancy
}
