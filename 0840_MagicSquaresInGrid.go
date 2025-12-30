/*
MATH THEORY: 
Based on given constraints of magic square,
(1) Centre of magic square must be 5, 
(2) sum must be 15.
(1) -> (3) The 5 cannot be at the edge (i.e. index cannot be 0 or row/col for row/col)
(4) Magic squares are symmetrical under rotation and reflection about the centre.

IDEA (BRUTE FORCE):
(1) Loop i from 1 to row, j from 1 to col to locate 5s
(2) When 5 is detected, check l=r, ul=br, u=b, ur=bl. If yes, increment by 1
(3) NEED TO CHECK IF prev row and prev col add up to 15 (opp will auto add up due to math)
(4) 5 cannot be at edge
(5) cannot have duplicates (I think this is already handled by other criteria)
*/

func numMagicSquaresInside(grid [][]int) int {
    //initialize
    var fout int
    row := len(grid)
    col := len(grid[0])
    if row < 3 || col < 3 { //too small
        fmt.Println("GRID TOO SMALL")
        return 0
    }
    var b int

    for i := 1; i < row - 1; i++{
        for j := 1; j < col - 1; j++ {
            if grid[i][j] == 5 {
                fmt.Printf("core detected at row %d col %d\n",i,j)
                b = 1
                b &= eye(grid[i][j-1],grid[i][j+1])
                b &= eye(grid[i-1][j-1],grid[i+1][j+1])
                b &= eye(grid[i-1][j],grid[i+1][j])
                b &= eye(grid[i-1][j+1],grid[i+1][j-1])
                b &= tri(grid[i-1][j-1],grid[i-1][j],grid[i-1][j+1])
                b &= tri(grid[i-1][j-1],grid[i][j-1],grid[i+1][j-1])
                fout += b
            }
        }
    }
    return fout
}

func eye(x, y int) int { //Use this to verify magic square
    if x*y == 0 || x + y != 10 || (x*y)%5 == 0 { //magic square cannot contain zero
        return 0
    }
    return 1
}

func tri(x, y, z int) int {
    if x + y + z == 15 {
        return 1
    }
    return 0
}
