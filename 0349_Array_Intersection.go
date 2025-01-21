func intersection(nums1 []int, nums2 []int) []int {
    count1 := make(map[int]int)
    count2 := make(map[int]int)
    fout := make([]int,0)
    for _,v := range(nums1){
        count1[v]++
    }
    for _,v := range(nums2){
        count2[v]++
    }
    for key := range(count1){
        _,ok := count2[key]
        if ok{
            fout = append(fout,key)
        }
    }
    return fout
}
