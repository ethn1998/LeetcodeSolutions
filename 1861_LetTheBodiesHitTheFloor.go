/*
IDEA:
(1) push everything to right
(2) transpose matrix
(3) reflect horizontally
NOTES:
byte val
'.' : 46
'#' : 35
'*' : 42
*/
func rotateTheBox(boxGrid [][]byte) [][]byte {
    m := len(boxGrid)
    n := len(boxGrid[0])
    for _,row := range boxGrid {
        bot := n-1
        for i := n-1; i >= 0; i-- {
            if row[i] == '*' {
                bot = i-1
            } else {
                if row[i] == '#' {
                    row[i] = '.' //remove body
                    row[bot] = '#' //put at bottom
                    bot-- //next body can only go on top of previous
                }
            }
        }
    }
    //fmt.Println(boxGrid) //DEBUG
    fout := make([][]byte,n)
    for i := 0; i < n; i++ {
        fout[i] = make([]byte,m)
        for j := 0; j < m; j++ {
            fout[i][j] = boxGrid[j][i]
        }
        for j := 0; j < m/2; j++ {
            //fmt.Println("DEBUG:",i,j)
            x := fout[i][j]
            fout[i][j] = fout[i][m-1-j]
            fout[i][m-1-j] = x
        }
    }
    //fmt.Println(fout) //DEBUG

    return fout
}
