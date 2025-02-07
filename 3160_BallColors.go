func queryResults(limit int, queries [][]int) []int {
    //Crude algorithm: Simulate process and count colors using a hashmap after each query.
    var ball, col int //Variables for parsing query
    ncolors := 0 //Variable for counting number of colors after each query
    //balls := make([]int,limit+1) //Simulate limit + 1 balls. This leads to MLE runtime error
    balls := make(map[int]int) //Spawn balls as they get colored. REMARK: We don't need the limit to implement this.
    colorcount := make(map[int]int) //Hashmap for counting balls of each color.
    result := make([]int,0) //Output
    for _,q := range(queries){
        //ncolors := 0 DO NOT RE-COUNT ON EACH STAGE.
        //Parse query
        ball = q[0] //Index of target ball
        col = q[1] //Color of target ball

        if balls[ball] != 0 { //If ball is already colored
            colorcount[balls[ball]] -= 1 //Uncolor ball
            if colorcount[balls[ball]] == 0 {
                ncolors--
            }
        }

        //Color ball with new color
        balls[ball] = col
        colorcount[col]++
        if colorcount[col] == 1{ //Only possible at this stage if coloring an uncolored ball
            ncolors++
        }

        /* DO NOT RE-COUNT ON EACH STAGE
        for _,val := range(colorcount){ //Range over colors
            if val != 0 { //This line is a bit clunky
                ncolors++
            }
        }
        */
        result = append(result,ncolors)
    }
    return result
}
