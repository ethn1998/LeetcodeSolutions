func judgeCircle(moves string) bool {
    var x,y int
    for _,char := range(moves){
        switch char{
            case 'R': x++
            case 'U': y++
            case 'L': x--
            case 'D': y--
        }
    }
    return (x == 0) && (y == 0)
}
