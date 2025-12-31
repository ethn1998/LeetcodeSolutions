func latestDayToCross(row int, col int, cells [][]int) int { //can get last day via binary search
    var mid, t, i, j int
    left := col - 1
    right := row * col
    board := make([][]int, row)
    for k := range board{
        board[k] = make([]int,col)
    }
    for left < right {
        //fmt.Println(t,board)
        mid = left + (right - left)/2
        if mid > t { // add water
            for k := t; k < mid; k++ {
                cell := cells[k] //adj 1-indexed
                fmt.Println(k, "water", cell)
                i = cell[0]-1
                j = cell[1]-1
                //fmt.Println(i,j)
                board[i][j] = 1
                fmt.Println(board)
            }
        } else if mid < t { // add land
            for k := t; k >= mid; k-- {
                cell := cells[k]
                fmt.Println(k, "land", cell)
                i = cell[0]-1
                j = cell[1]-1
                //fmt.Println(i,j)
                board[i][j] = 0
                fmt.Println(board)
            }
        } else {
            break
        }
        //fmt.Println(mid,board)
        fmt.Println(t,mid,board)
        t = mid
        //BFS to check if we can cross
        if bfs(board) {
            //fmt.Println("Can Cross")
            left = mid
        } else {
            //fmt.Println("Cannot Cross")
            right = mid - 1
        }
        fmt.Println("DEBUG BOARD",board)
        fmt.Println("update leftright",left, right)
    }
    return left
}

func bfs(cells [][]int) bool{//BFS to check if we can cross
    //fmt.Println("bfsin:",cells)
    var start bool
    nrow := len(cells)
    ncol := len(cells[0])

    //fmt.Println("lims:",nrow,ncol)
    var row, col, index, index1 int

    mem := make(map[int]bool) //use memory of vectorized indices of visited cells
    for n := 0; n < ncol; n++ {
        //fmt.Println(start,n,cells[0][n])
        if !start && cells[0][n] == 0 {
            //fmt.Println("Start col:",n)
            queue := []int{n}
            mem[n] = true
            start = true
            for k := 0; k < len(queue); k++ {
                index = queue[k]
                row = index / ncol
                col = index % ncol

                //cells[row][col] = 1
                //down
                if cells[row+1][col] == 0 && !mem[col + ncol*(row + 1)] {
                    if row == nrow - 2 {
                        return true
                    } else {
                        index1 = col + (row + 1)*ncol
                        mem[index1] = true
                        queue = append(queue,index1)
                    }
                }
                if col < ncol - 1 { //right
                    if cells[row][col+1] == 0  && !mem[col + 1 + row*ncol]{
                        index1 = col + 1 + row*ncol
                        mem[index1] = true
                        queue = append(queue,index1)
                    }
                }
                if col > 0 { //left
                    if cells[row][col-1] == 0 && !mem[col - 1 + row*ncol] {
                        index1 = col - 1 + row*ncol
                        mem[index1] = true
                        queue = append(queue,index1)
                    }    
                }
                if row > 0 { //up
                    if cells[row-1][col] == 0 && !mem[col + (row - 1)*ncol] {
                        index1 = col + (row - 1)*ncol
                        mem[index1] = true
                        queue = append(queue,index1)
                    }
                }
            }
        }
        if start && cells[0][n] == 1 {
            start = false
        }
    }    
    return false
}
