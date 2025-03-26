func minOperations(grid [][]int, x int) int {
	var fout int
	arr := make([]int, len(grid)*len(grid[0]))
	//Only possible to equalize if g_{ij} mod x are equal for all i,j
	r := grid[0][0] % x
	for i, row := range grid {
		for j, v := range row {
			if v%x != r {
				return -1
			}
			arr[len(grid[0])*i+j] = v - r
		}
	}
	//Find number of operations to equalize.
	//Mean? Median? Mode?
	//Hint 3: sort suggests use of median
	//Why is median the best?
	sort.Ints(arr)
	if len(arr) != 0 {
		m := arr[len(arr)/2] //Median
		for _, v := range arr {
			if m > v {
				fout += m - v
			} else {
				fout += v - m
			}
		}
	} else { //If array has even length
		var fout0, fout1 int
		m0 := arr[len(arr)/2]
		m1 := arr[len(arr)/2]
		for _, v := range arr {
			if m0 > v {
				fout0 += m0 - v
			} else {
				fout0 += v - m0
			}
			if m1 > v {
				fout1 += m1 - v
			} else {
				fout1 += v - m1
			}
		}
		if fout0 > fout1 {
			fout = fout1
		} else {
			fout = fout0
		}
	}
	return fout/x
}
