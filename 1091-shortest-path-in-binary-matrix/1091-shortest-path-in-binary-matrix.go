func shortestPathBinaryMatrix(grid [][]int) int {
    if grid[0][0] == 1 {
        return -1
    }
    numR, numC := len(grid), len(grid[0])
    dirs := [8][2]int{
        {-1,0}, // N
         {1,0}, // S
         {0,1}, // E
         {0,-1}, // W
         {-1,1}, // NW
         {-1,-1}, // NE
         {1,1},  // SW
         {1,-1}, // NE
    }
    var (
        queue [][2]int
        currNode [2]int
    )
    visited := make([][]bool, numR)
    for x := range visited {
        visited[x] = make([]bool, numC)
    }
    visited[0][0] = true
    queue = append(queue, [2]int{0,0})
    shortestDist := 1
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            currNode, queue = queue[0], queue[1:]
            x, y := currNode[0], currNode[1]
            if x == numR-1 && y == numC-1 {
                return shortestDist
            }
            for _, dir := range dirs {
                dX, dY := dir[0], dir[1]
                newX, newY := x + dX, y + dY 
                if newX < 0 || newX >= numR || newY < 0 || newY >= numC {
                    continue
                }
                if grid[newX][newY] == 1 || visited[newX][newY] {
                    continue
                }
                visited[newX][newY] = true
                queue = append(queue, [2]int{newX, newY})
            }
        }
        shortestDist++
    }
    return -1
}