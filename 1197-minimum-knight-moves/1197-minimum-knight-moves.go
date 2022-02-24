
// 1197. Minimum Knight Moves


// bfs approach
// time: O((max(|x|, |y|))^2) since cells we cover has an upper bound 
// roughly the area of the rectangle inscribed in this circle 

// O(1) space for visited set since size independent of input
// but queue size grows depending on how far the (x, y) input is from the origin 
// space: O((max(|x|, |y|))^2) since cells we cover has an upper bound 
// roughly the area of the rectangle inscribed in this circle 

// constraint: -300 <= x, y <= 300
// 0 <= |x| + |y| <= 300
// var visited [601][601]bool  
// (i, j) = (0,0) is located @ visited[i+300][i+300] 
// e.g. (-2, -1) is located @ visited[298][299] 
// e.g. (300, 300) is located @ visited[600][600]

func minKnightMoves(x int, y int) int {
    dX := [8]int{2, 2, -2, -2, 1, 1, -1, -1}
    dY := [8]int{1, -1, 1, -1, 2, -2, 2, -2}
    var (
        dist int
        r, c int
        visited [601][601]bool
        node [2]int
        queue [][2]int
    )
    visited[0][0] = true
    queue = append(queue, [2]int{0,0})
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            r, c = node[0], node[1]
            if r == x && c == y {
                return dist
            }
            for idx := 0; idx < 8; idx++ {
                nextR, nextC := dX[idx] + r, dY[idx] + c
                if nextR < -300 || nextR > 300 || nextC < -300 || nextC > 300 {
                    continue
                }
                if visited[nextR+300][nextC+300] {
                    continue
                }
                queue = append(queue, [2]int{nextR, nextC})
                visited[nextR+300][nextC+300] = true
            }
        }
        dist++
    }
    return -1
}