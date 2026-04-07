type Pos struct {
    x int
    y int
}

func newPos() *Pos {
    return &Pos{0,0}
}

func (p Pos) getSqDist() int {
    return p.x * p.x + p.y * p.y
}

func robotSim(commands []int, obstacles [][]int) int {
    var fout int
    var x1,y1 int
    obsSet := make(map[Pos]bool)
    for _,obs := range obstacles {
        px := obs[0]
        py := obs[1]
        obsSet[Pos{px,py}] = true
    }
    //fmt.Println(obsSet)
    loc := newPos()
    dir := newPos()
    dir.y = 1 //start facing north
    //fmt.Println(dir)

    for _,k := range commands {
        if k > 0 {
            for i := 0; i < k; i++ {
                x1 = loc.x + dir.x
                y1 = loc.y + dir.y
                //fmt.Println(i,k,Pos{x1,y1},obsSet[Pos{x1,y1}])
                if obsSet[Pos{x1,y1}] { //collision detection
                    break
                }
                loc.x = x1
                loc.y = y1
            }
        } else if k == -2 { //rotate anticlockwise pi/2 rads
            x1 = -dir.y
            y1 = dir.x
            dir.x = x1
            dir.y = y1
        } else if k == -1 { //rotate clockwise pi/2 rads
            x1 = dir.y
            y1 = -dir.x
            dir.x = x1
            dir.y = y1
        }
        //fmt.Println(k,loc,dir)
        if fout < loc.getSqDist() {
            fout = loc.getSqDist()
        }
    }
    return fout
}
