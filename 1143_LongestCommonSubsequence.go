/*
HINT 2:
if text1[i] == text2[j] {
    DP[i][j] = DP[i - 1][j - 1] + 1
} else { //i.e. text1[i] != text2[j]
    DP[i][j] = max(DP[i - 1][j], DP[i][j - 1])
}

IDEA:
Recursively backtrack from (i,j) = (len(text1)-1, len(text2)-1)
Set initial condition for (i,j) = (0,0)

Clearly, max answer = min(len(text1),len(text2))

TLE on even simpler case.
Model solutions for this kind of problems define a DP array
We can use a hashmap for this kind of functionality.
Existence of key to tell if cell has been visited before.
*/

func longestCommonSubsequence(text1 string, text2 string) int {
    //fmt.Println(len(text1),len(text2))
    //encode i = n % p, j = n / p, max n = p*len(text2) - 1

    var fout int
    //fout = dp(text1, text2, len(text1)-1, len(text2)-1)
    mem := make(map[int]int) 
    fout = dp(text1, text2, len(text1)*len(text2)-1, mem)

	return fout
}

func dp(text1, text2 string, key int, mem map[int]int) int {
	var fout int
    val, ok := mem[key]
    //fmt.Println("mem:",key,val,ok)
    if ok {
        return val
    }
    p := len(text1)
    i := key % p
    j := key / p

    //fmt.Println("In:",i,j,text1[i],text2[j])
	if i > 0 && j > 0 {
		if text1[i] == text2[j] {
			fout = dp(text1, text2, (i-1) + (j-1)*len(text1), mem) + 1
		} else {
			y1 := dp(text1, text2, i-1 + j*len(text1), mem)
			y2 := dp(text1, text2, i + (j-1)*len(text1), mem)
			fout = max(y1, y2)
		}
	} else { //boundary where i == 0 || j == 0
        if j > 0 {
            for k := 0; k <= j; k++ {
                if text1[0] == text2[k] {
                    fout = 1
                    break
                }
            }
        } else if i > 0 {
            for k := 0; k <= i; k++ {
                if text1[k] == text2[0]{
                    fout = 1
                    break
                }
            }
        } else { //case where i == j == 0
            if text1[0] == text2[0] {
                fout = 1
            }
        }
    }
    mem[key] = fout
	return fout
}
