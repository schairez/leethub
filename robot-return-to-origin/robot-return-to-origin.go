type Point struct {
    x, y int
}
/*
1 <= moves.length <= 2 * 104
moves only contains the characters 'U', 'D', 'L' and 'R'.
*/

func judgeCircle(moves string) bool {
    start := &Point{0,0} 
    for _, move := range []byte(moves) {
        switch move {
        case 'U':
            start.y++
        case 'D':
            start.y--
        case 'L':
            start.x--
        case 'R':
            start.x++
        }
    } 
    return start.x == 0 && start.y == 0
}
/*
"LDRRLRUULR"
*/