
type invl struct { //record intervals
    l int
    r int
    sum int
}

func maxSumTrionic(nums []int) int64 {
    minf := -1 << 63
    var l,sum int
    var desc bool
    n := len(nums)
    descs := make([]*invl,0)
    for i := 1; i < n - 1; i++ { //defn: desc part of trionic arrays must start from i >= 1
        if desc {
            sum += nums[i]
            if nums[i] <= nums[i+1] {
                desc = false
                descs = append(descs, &invl{l,i,sum})
            }
        } else {
            if nums[i] > nums[i+1] {
                desc = true
                sum = nums[i]
                l = i
            }
        }
    }
    //fmt.Println(len(descs)) //DEBUG
    //NOTE: If nums is still desc at end, don't bother because the desc interval would not form part of a trionic array anyways.
    fout := minf
    //lsums := make([]int,0)
    //rsums := make([]int,0)

    for _,iv := range descs {
        var lsum, rsum int
        l0 := iv.l
        l := l0 - 1
        r0 := iv.r
        r := r0 + 1
        //flush
        lsums := make([]int,0)
        rsums := make([]int,0)
        for l >= 0 {
            if nums[l] < nums[l+1] {
                lsum += nums[l]
                lsums = append(lsums,lsum)
            } else {
                break
            }
            l--
        }
        for _,v := range lsums {
            lsum = max(v,lsum)
        }
        for r < n {
            if nums[r-1] < nums[r] {
                rsum += nums[r]
                rsums = append(rsums,rsum)
            } else {
                break
            }
            r++
        }
        for _,v := range rsums {
            rsum = max(v,rsum)
        }
        if len(lsums) > 0 && len(rsums) > 0 {
            fout = max(fout, lsum + iv.sum + rsum)
        }
        /*DEBUG
        fmt.Println(iv)
        fmt.Println(lsums,rsums)
        fmt.Printf("lsum: %d msum: %d rsum: %d\n",lsum,iv.sum,rsum)
        */
    }
    return int64(fout)
}
