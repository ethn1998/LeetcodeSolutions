/*
Operation preserves product parity (odd or even number of -ve elements) of matrix.
Two cases:
(1) Even: Can eventually turn everything +ve: Answer is just the sum of abs of all elements
(2) Odd: All but one are +ve.

Special case: zero. In this case, we can always set everything to +ve by freely introducing singular flips using the zero
*/

func maxMatrixSum(matrix [][]int) int64 {
    var fout, val int64
    var n bool //XOR sum of int sign bit; coerce to bool with 0: false, 1: true
    var z bool //to track if zero exists
    eps := int64(math.MaxInt64) //smallest non-zero seen
    for _, row := range matrix{
        for _, val0 := range row {
            val = int64(val0)
            if val > 0 {
                fout += val
                if val < eps {
                    eps = val
                }
            } else if val < 0 {
                fout -= val
                n = !n
                if -val < eps {
                    eps = -val
                }
            } else { //zero
                z = true
            }   
        }
    }
    if z || !n {
        return fout
    }
    return fout - 2*eps
}
