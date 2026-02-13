//NOTE: This does NOT scale well if s has more characters than just a,b,c.

type pair struct {
    ab int //a-b
    ac int //a-c
}

func longestBalanced(s string) int {
    var fout,a,b,c int
    //case: 1 letter
    streak := 1
    //case: 2 letter
    var mpab = map[int]int{0:-1}
    var mpac = map[int]int{0:-1}
    var mpbc = map[int]int{0:-1}

    //case: 3 letter
    var mpabc = map[pair]int{{0,0}:-1}
    
    for i, char := range s {
        //update
        if char == 'a' {
            mpbc = make(map[int]int) //clear
            a++
        } else if char == 'b' {
            mpac = make(map[int]int) //clear
            b++
        } else { //char = 'c'
            mpab = make(map[int]int) //clear
            c++
        }
        //case: 1 letter
        if i > 0 {
            if byte(char) == s[i-1] {
                streak++
            } else {
                streak = 1
            }
        }
        fout = max(streak,fout)
        //case: 2 letter
        j, ok := mpab[a-b]
        if ok {
            fout = max(fout, i-j)
        } else {
            mpab[a-b] = i
        }

        j, ok = mpac[a-c]
        if ok {
            fout = max(fout, i-j)
        } else {
            mpac[a-c] = i
        }

        j, ok = mpbc[b-c]
        if ok {
            fout = max(fout, i-j)
        } else {
            mpbc[b-c] = i
        }

        //case: 3 letter
        j, ok = mpabc[pair{a-b,a-c}]
        if ok {
            fout = max(fout,i-j)
        } else {
            mpabc[pair{a-b,a-c}] = i
        }
    }
    return fout
}
