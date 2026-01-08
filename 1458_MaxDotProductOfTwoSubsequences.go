//Remark: Recursion to prevent calculating unnecessary stuff!

func maxDotProduct(nums1 []int, nums2 []int) int {
    var fout int
    mem := make(map[int]int)
    fout = dp(nums1, nums2, len(nums1)*len(nums2)-1, mem)
    return fout
}

func dp(nums1, nums2 []int, key int, mem map[int]int) int {
    //var fout int
    val, ok := mem[key]
    if ok {
        return val
    }

    i := key % len(nums1)
    j := key / len(nums1)
    //fmt.Println("In:",i,j,mem)

    fout := nums1[i] * nums2[j]
    if i == 0 || j == 0 {
        if i > 0 {
            fout = getMaxProduct(nums2[0],nums1[0:i+1])
        } else if j > 0 {
            fout = getMaxProduct(nums1[0],nums2[0:j+1])
        } else { //case i = 0, j = 0
            fout = nums1[0] * nums2[0]
        }
    } else {
        if dp(nums1, nums2, (i-1) + (j-1)*len(nums1), mem) > 0 {
            fout += dp(nums1, nums2,i-1 + (j-1)*len(nums1), mem)
        }
        if fout < dp(nums1, nums2, i-1 + (j)*len(nums1), mem) {
            fout = dp(nums1, nums2, i-1 + (j)*len(nums1), mem)
        }
        if fout < dp(nums1, nums2, i + (j-1)*len(nums1), mem) {
            fout = dp(nums1, nums2, i + (j-1)*len(nums1), mem)
        }
    }
    mem[key] = fout
    return fout
}


func getMaxProduct(n int, vec []int) int {
    //fmt.Println("MaxProductIn:",n,vec)
    if n == 0 || len(vec) == 0 {
        return 0
    }
    fout := -1 << 63
    for _, v := range vec {
        if n * v > fout {
            fout = n * v
        }
    }
    return fout
}
