/*
Rectangles are defined as []int{x1,y1,x2,y2}
CODE TRANSLATED FROM JAVA IMPLEMENTATION FOR LINE SWEEP
*/

func rectangleArea(rectangles [][]int) int {
    const OPEN = 0
    const CLOSE = 1
    events := make([][]int,2*len(rectangles))
    for i, rect := range rectangles {
        events[2*i] = []int{rect[1], OPEN, rect[0], rect[2]}
        events[2*i+1] = []int{rect[3], CLOSE, rect[0], rect[2]}
    }
    sort.Slice(events, func(i,j int) bool{
        return events[i][0] < events[j][0]
    })

    active := make([][]int,0)
    cur_y := events[0][0]
    var fout int
    for _,event := range events {
        y := event[0]
        typ := event[1]
        x1 := event[2]
        x2 := event[3]
        
        //calculate query, the length of the active breadth
        query := 0
        cur := -1
        for _, xs := range active {
            cur = max(cur,xs[0])
            query += max(xs[1]-cur,0)
            cur = max(cur,xs[1])
        }
        fout += query * (y - cur_y) //breadth * height
        if typ == OPEN { //add to active queries
            active = append(active,[]int{x1,x2})
            sort.Slice(active, func(i,j int) bool {
                return active[i][0] < active[j][0]
            })
        } else {
            for i := 0; i < len(active); i++ {
                if active[i][0] == x1 && active[i][1] == x2 {
                    active = append(active[:i],active[i+1:]...)
                }
            }
        }
        cur_y = y
    }
    return fout % (1e9+7)
}
