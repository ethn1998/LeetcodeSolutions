/*
Consider two intervals I0 = (l0,r0), I1 = (l1,r1)
TEST FOR INTERSECTION, CASE BY CASE
1: l0 < l1, then l1 < r0
2: l1 < l0, then l0 < r1 by symmetry??
3: l0 == l1, then we will always intersect
*/

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
    var fout int64
    var b, h int
    var l int64

    for i := 0; i < len(bottomLeft); i++ {
        for j := i + 1; j < len(bottomLeft); j++ {
            //horiz check
            b = getOverlap(bottomLeft[i][0],topRight[i][0],bottomLeft[j][0],topRight[j][0])
            //vert check
            h = getOverlap(bottomLeft[i][1],topRight[i][1],bottomLeft[j][1],topRight[j][1])
            if b > 0 && h > 0 {
                l = int64(min(b,h))
                if fout < l {
                    fout = l
                }
            }
        }
    }
    return fout * fout
}

func getOverlap(l0, r0, l1, r1 int) int {
    var fout int
    //fmt.Printf("In l0: %d r0: %d, l1: %d, r1: %d\n",l0,r0,l1,r1)
    if l0 < l1 {
        fout = min(r0 - l1,r0 - l0,r1 - l1)
    } else if l1 < l0 {
        fout = min(r1 - l0,r0 - l0,r1 - l1)
    } else {
        fout = min(r1,r0) - l0
    }
    //fmt.Println("Out:",fout)
    return fout
}
