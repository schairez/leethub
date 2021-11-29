//time: O(n)
//space: O(1); pt x and y vars are modified in place

type Point struct {
    x, y int
}
/*
1 <= moves.length <= 2 * 104
moves only contains the characters 'U', 'D', 'L' and 'R'.
*/

func judgeCircle(moves string) bool {
    cell := &Point{0,0} 
    for _, move := range []byte(moves) {
        switch move {
        case 'U':
            cell.y++
        case 'D':
            cell.y--
        case 'L':
            cell.x--
        case 'R':
            cell.x++
        }
    } 
    return cell.x == 0 && cell.y == 0
}
/*
"LDRRLRUULR"
*/