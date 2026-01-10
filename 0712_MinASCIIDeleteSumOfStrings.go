/*
IDEA:
The solution results in the longest common substring (LCS).
The common substring must have its ASCII sum MAXIMIZED. (Can get this using DP)
return sum(s1) + sum(s2) - 2*sum(LCS)
NOTE: If no common substring, delete EVERYTHING.
*/

func minimumDeleteSum(s1 string, s2 string) int {
    mem := make(map[int]int) //encryption: i : %, j : /
    x1 := getASCIISum(s1)
    x2 := getASCIISum(s2)

    x := dp(s1, s2, len(s1)*len(s2)-1, mem)

    return x1 + x2 - 2*x
}

func dp(s1, s2 string, key int, mem map[int]int) int { //substring with max ASCII sum
    val, ok := mem[key]
    if ok {
        return val
    }
    var fout int
    i := key % len(s1)
    j := key / len(s1)
    a1 := int(s1[i])
    a2 := int(s2[j])
    if i > 0 && j > 0 {
        if a1 == a2 {
            fout = a1 + dp(s1,s2,(i-1)+(j-1)*len(s1),mem)
        } else {
            x1 := dp(s1,s2,(i)+(j-1)*len(s1),mem)
            x2 := dp(s1,s2,(i-1)+(j)*len(s1),mem)
            fout = max(x1,x2)
        }
    } else {
        if i > 0 {
            for k := 0; k <= i; k++ {
                if s1[k] == s2[0] {
                    fout = a2
                    break
                }
            }
        } else if j > 0 {
            for k := 0; k <= j; k++ {
                if s1[0] == s2[k] {
                    fout = a1
                    break
                }
            }
        } else {
            if a1 == a2 {
                fout = a1
            }
        }
    }
    mem[key] = fout
    return fout
}

func getASCIISum(s string) int {
    var fout int
    for _,char := range s {
        fout += int(char)
    }
    return fout
}
