import "sort"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    //These arrays are actually hashmaps
    mapout := make(map[int]int)
    for _,p := range(nums1){
        mapout[p[0]] += p[1]
    }
    for _,p := range(nums2){
        mapout[p[0]] += p[1]
    }
    fout := make([][]int,0)
    for key,val := range(mapout){
        fout = append(fout,[]int{key,val})
    }
    sort.Slice(fout, func(i,j int) bool{
        return fout[i][0] < fout[j][0]
    })
    return fout
}
