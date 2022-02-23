

func minKnightMoves(x int, y int) int {
    visited := make(map[int]map[int]bool, 300)
    dX := [8]int{2, 2, -2, -2, 1, -1, 1, -1}
    dY := [8]int{1, -1, 1, -1, 2, 2, -2, -2}
    numMoves := 0
    // 8 possible next moves
    visited[0] = make(map[int]bool)
    visited[0][0] = true
    var queue [][2]int //0 -> x, 1 -> y
    queue = append(queue, [2]int{0,0})
    var node [2]int 
    var nextX, nextY int
    for len(queue) != 0  {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            nextX, nextY = node[0], node[1]
            if nextX == x && nextY == y {
                return numMoves
            }
            for idx := 0; idx < 8; idx++ {
                nextX, nextY := node[0] + dX[idx], node[1] + dY[idx]
                if visX, existsX := visited[nextX]; existsX {
                    if _, existsY := visX[nextY]; existsY {
                        continue
                    }
                } else {
                    visited[nextX] = make(map[int]bool)
                }
                visited[nextX][nextY] = true 
                queue = append(queue, [2]int{nextX, nextY})
            }
        }
        numMoves++
    }
    return -1
}