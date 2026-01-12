func minTimeToVisitAllPoints(points [][]int) int {
    var fout,dx,dy int
    x := points[0][0]
    y := points[0][1]
    for _, point := range points {
        dx =  point[0] - x
        if dx < 0 {
            dx = -dx
        }
        dy = point[1] - y
        if dy < 0 {
            dy = -dy
        }
        if dx > dy {
            fout += dx
        } else {
            fout += dy
        }
        x = point[0]
        y = point[1]
    }
    return fout

}
