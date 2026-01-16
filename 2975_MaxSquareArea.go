func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
    fout := -1
    hFences = append(hFences,1,m)
    vFences = append(vFences,1,n)
    sort.Slice(hFences, func(i,j int) bool {
        return hFences[i] < hFences[j]
    })
    sort.Slice(vFences, func(i,j int) bool {
        return vFences[i] < vFences[j]
    })

    hmem := make(map[int]bool) //set of possible heights
    bmem := make(map[int]bool) //set of possible breadths

    for i := 0; i < len(hFences); i++ {
        for j := i+1; j < len(hFences); j++{
            hmem[hFences[j] - hFences[i]] = true
        }
    }
    for i := 0; i < len(vFences); i++ {
        for j := i+1; j < len(vFences); j++{
            bmem[vFences[j] - vFences[i]] = true
        }
    }

    hvec := make([]int,0)
    bvec := make([]int,0)

    for key := range hmem {
        hvec = append(hvec,key) 
    }
    for key := range bmem {
        bvec = append(bvec,key) 
    }

    sort.Slice(hvec,func(i,j int) bool {
        return hvec[i] > hvec[j]
    })
    sort.Slice(bvec,func(i,j int) bool {
        return bvec[i] > bvec[j]
    })

    var i,j int

    for i < len(hvec) && j < len(bvec){
        //fmt.Println(i,j,hvec[i],bvec[j],hvec[i]*bvec[j])
        if hvec[i] == bvec[j] {
            return (hvec[i] * bvec[j]) % (1e9+7) 
        } else if hvec[i] > bvec[j] {
            i++
        } else {
            j++
        }
    }

    return fout
}
