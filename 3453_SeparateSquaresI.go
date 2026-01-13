/*
Try binary search.
NOTE: x seems redundant
*/

func separateSquares(squares [][]int) float64 {
    //init using problem constraints
    left := 1.0e9
    right := 0.0
    var y,l,h,totarea float64
    for _, square := range(squares){
        y = float64(square[1])
        l = float64(square[2])
        if y < left {
            left = y
        }
        if y + l > right {
            right = y + l
        }
        totarea += l * l
    }
    for left < right { //for floats, need to specify convergence criterion
        mid := left + (right - left) * 0.5
        t := 0.0 //area below mid line
        for _,square := range(squares) {
            y = float64(square[1])
            if y < mid {
                l = float64(square[2])
                h = math.Min(l, mid - y)
                t += 2.0 * h * l
            }
        }
        if t < totarea {
            left = mid
        } else {
            right = mid
        } 
        if right - left < 1.0e-5 {
            break
        }
    }
    return left
}
