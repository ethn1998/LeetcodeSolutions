type Robot struct {
    xlim int
    ylim int
    idx int
    pst bool //persistent boolean to check if robot has moved before
}


func Constructor(width int, height int) Robot {
    return Robot{width,height,0,false}
}


func (this *Robot) Step(num int)  {
    this.pst = true
    this.idx += num
    this.idx = this.idx % (2 * (this.xlim-1) + 2 * (this.ylim-1))
}


func (this *Robot) GetPos() []int {
    var x,y int
    if this.idx < this.xlim {
        x = this.idx
        //y = 0
    } else if this.idx < this.xlim -1 + this.ylim {
        x = this.xlim - 1
        y = this.idx - (this.xlim - 1)
    } else if this.idx < this.xlim -1 + this.ylim -1 + this.xlim {
        x = this.xlim - 1 - (this.idx -  (this.xlim -1 + this.ylim -1))
        y = this.ylim - 1
    } else {
        //x = 0
        y = this.ylim - 1 - (this.idx - (this.xlim -1 + this.ylim -1 + this.xlim -1))
    }
    return []int{x,y}
}


func (this *Robot) GetDir() string {
    //IDEA: robot can only walk along perimeter -> direction solely depends on the edge robot is on.
    if this.pst {
        if this.idx == 0 {
            return "South"
        } else if this.idx < this.xlim {
            return "East"
        } else if this.idx < this.xlim - 1 + this.ylim {
            return "North"
        } else if this.idx < this.xlim - 1 + this.ylim - 1 + this.xlim {
            return "West"
        } else {
            return "South"
        }
    }
    return "East"
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Step(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
