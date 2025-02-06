func tupleSameProduct(nums []int) int {
    /*
    REMARK: This looks like a more complex version of two sum.
    IDEA: Each combination abcd contributes 8 tuples. (See example 1.)
    */
    // IDEA: Try following hints first?? It gets the job done, but O(n^2) complexity looks bleak.
    var fout int
    plist := make(map[int]int) //Count number of times product is formed
    for i,u := range(nums){
        for _,v := range(nums[i+1:]){
            plist[u*v]++
        }
    }
    for _,v := range(plist){
        fout += 4*v*(v-1)
    }
    return fout
}
