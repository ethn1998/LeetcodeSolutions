func twoSum(nums []int, target int) []int {
    var fout []int
    //var trigger bool
    mem := make(map[int]int) //Input value output index
    for i,v := range nums{
        j, ok := mem[target-v] //Check if value has been seen before
        if ok {
            fout = append(fout,i,j)
            break
        }
        mem[v] = i
    }
    return(fout)
}
