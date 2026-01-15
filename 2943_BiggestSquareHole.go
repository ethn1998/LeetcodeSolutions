/*
why not just remove all removable bars and find biggest square subregion?
b = max consecutive subsequence of vBars
h = max consecutive subsequence of hBars
l = min(b,h)
return l*l
NOTE: problem constraints suggest that we can never remove outer boundary
*/

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	//sorting is probably redundant, but ensures that the pointer method works properly later
	sort.Slice(hBars, func(i, j int) bool {
		return hBars[i] < hBars[j]
	})
	sort.Slice(vBars, func(i, j int) bool {
		return vBars[i] < vBars[j]
	})
	var b, h int //init for final answer
	//loop through horizontal bars
    //Two pointer method to find longest consecutive subsequence
    for i := 0; i < len(hBars); {
        j := i
        for j < len(hBars)-1 {
            if hBars[j] != hBars[j+1] - 1{
                break
            }
            j++
        }
        if h < j - i + 1 {
            h = j - i + 1
        }
        i = j + 1
    }

    for i := 0; i < len(vBars); {
        j := i
        for j < len(vBars)-1 {
            if vBars[j] != vBars[j+1] - 1{
                break
            }
            j++
        }
        if b < j - i + 1 {
            b = j - i + 1
        }
        i = j + 1
    }

	//calculate area of biggest square
    l := min(b, h) + 1
	return l * l
}
