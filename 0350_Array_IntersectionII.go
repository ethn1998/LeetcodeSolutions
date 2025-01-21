func intersect(nums1 []int, nums2 []int) []int {
    var n,Omega int
    count1 := make(map[int]int)
    count2 := make(map[int]int)
    fout := make([]int,0)
    for _,v := range(nums1){
        count1[v]++
    }
    for _,v := range(nums2){
        count2[v]++
    }
    for key1,val1 := range(count1){
        val2 := count2[key1]
        n = 0
        Omega = min(val1,val2)
        for n < Omega{
            fout = append(fout,key1)
            n++
        }
    }
    return fout
}
