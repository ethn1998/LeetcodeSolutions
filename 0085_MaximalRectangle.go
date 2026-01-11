/*
A rectangle can be defined as a region (x1,y1), (x2,y2) such that
(1) x1 <= x2, 
(2) y1 <= y2, 
(3) for all i,j s.t. x1 <= i <= x2, y1 <= j <= y2, mat[i][j] = 1
*/

func maximalRectangle(matrix [][]byte) int {
    mem := make(map[int]int)
    fout := dp(matrix, len(matrix)*len(matrix[0])-1, mem)
    return fout
}

func dp(mat [][]byte, key int, mem map[int]int) int {
    var fout int
    val, ok := mem[key]
    if ok {
        return val
    }
    n := len(mat)
    i := key % n
    j := key / n
    //fmt.Println("In:",key,n,i,j,mat[i][j]-48)
    if i > 0 || j > 0 {
        fout = 0
        b := j //largest rectangle containing bottom right element of array
        for h := 0; h <= i; h++ {
            //fmt.Println("i:",i,"trial h:",h)
            if mat[i-h][j] == 48 { //stop when you see zero
                break
            }
            left := j - b
            right := j
            for left < right {
                mid := left + (right - left) / 2
                //fmt.Println("binom h:", h,"l:",left,"m:",mid,"r:",right)
                updateleft := false
                for k := mid; k <= right; k++ {
                    //fmt.Println("row:",i-h,"col:",k,"val:",mat[i-h][k]-48)
                    if mat[i-h][k] == 48 {
                        updateleft = true
                        break
                    }
                }
                if updateleft {
                    left = mid + 1
                } else {
                    right = mid
                }
            }
            if b > j - right {
                b = j - right
            }
            //fmt.Println("h:",h,"binomout:",right,"b:",b)
            if (b + 1) * (h + 1) > fout {
                fout = (b + 1) * (h + 1) 
            }
            //fmt.Println("area:",fout)
        }
        if i > 0 {
            fout = max(fout, dp(mat,(i-1)+j*n,mem))
        }
        if j > 0 {
            fout = max(fout, dp(mat,(i)+(j-1)*n,mem))
        }
    } else { //case where i == j == 0
        fout = int(mat[0][0]) - 48
    }
    mem[key] = fout
    //fmt.Println("Out:",key,mem[key])
    return fout
}
