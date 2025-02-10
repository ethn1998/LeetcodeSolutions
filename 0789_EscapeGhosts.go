func escapeGhosts(ghosts [][]int, target []int) bool {
    /*
    Use Manhattan distance
    Lazy idea: Have ghosts also run to target and then do nothing.
    If ghost reaches target before you, they can just catch you when you reach the target
    */
    var d0, d, dx, dy int
    if target[0] > 0 {
        d0 += target[0]
    } else {
        d0 -= target[0]
    }
    if target[1] > 0 {
        d0 += target[1]
    } else {
        d0 -= target[1]
    }

    for _, ghost := range(ghosts){
        d = 0
        dx = ghost[0] - target[0]
        if dx > 0 {
            d += dx
        } else {
            d -= dx
        }
        dy = ghost[1] - target[1]
        if dy > 0 {
            d += dy
        } else {
            d -= dy
        }
        if d <= d0 {
            return false
        }
    }
    return true
}
