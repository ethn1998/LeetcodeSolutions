/*
row i has i+1 glasses, 0 <= i < 100
Let g(i,j) represent the jth glass of the ith row where 0 <= j <= n
If g(i,j) is full, excess flows EQUALLY to g(i+1,j) and g(i+1,j+1)
*/

func champagneTower(poured int, query_row int, query_glass int) float64 {
    //simulate solution?
    if query_row == 0 { //common sense case
        return math.Min(float64(poured),1.0)
    }
    //init from row = 0
    prev := []float64{float64(poured - 1)} //amount that will overflow into next row
    row := 1 //row number. Integer.
    for row <= query_row {
        var ctn bool
        curr := make([]float64,row + 1)
        for i, v := range prev {
            if v > 0 {
                curr[i] += 0.5 * v
                curr[i+1] += 0.5 * v
            }
        }
        if row < query_row {
            prev = make([]float64,len(curr))
            for i,v := range(curr) {
                if v > 1.0 {
                    prev[i] = v - 1.0
                    ctn = true
                }
            }
        } else {
            return math.Min(curr[query_glass],1.0)
        }
        if ctn {
            row++
        } else {
            break
        }
    }
    return 0.0
}
