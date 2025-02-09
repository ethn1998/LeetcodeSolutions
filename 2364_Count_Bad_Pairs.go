func countBadPairs(nums []int) int64 {
  //This naive implementation leads to TLE for large inputs.  
  var fout int64 //Initialize and declare output type.
    for i,u := range(nums){
        for d,v := range(nums[i+1:]){ 
            //In this implementation, d is the index on the new array, i.e. d+1 = the difference between indices.
            if d+1 != v-u {
                fout++
            }
        }
    }
    return fout
}


func countBadPairs(nums []int) int64 {
    /*
    IDEA: (HINT 2)
    j-i != v-u iff j-v != i-u -> Group into different i-u's, and then consider perms and combs.
    Pairs within the same key form non-bad pairs.
    For any group, there are n choose 2 = n!/((n-2)!2!) = n(n-1)/2 pairs
    */

    groups := make(map[int]int64)
    n := int64(len(nums))
    for i,u := range(nums){
        groups[i-u]++
    }
    fout = n*(n-1)/2
    for _,v := range(groups){
        fout -= v*(v-1)/2 //Remove good pairs
    }

    return fout
}
