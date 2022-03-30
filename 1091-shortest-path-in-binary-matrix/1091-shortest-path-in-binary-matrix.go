func shortestPathBinaryMatrix(grid [][]int) int {
    if grid[0][0] == 1 {
        return -1
    }
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
    numR, numC := len(grid), len(grid[0])
    visited := make([][]bool, numR)
    for x := range visited {
        visited[x] = make([]bool, numC)
    }
    srcNode := [2]int{0,0}
    x, y := srcNode[0], srcNode[1]
    var (
        queue [][2]int 
        currNode [2]int
        dist     int
    )
    queue = append(queue, srcNode)
    visited[x][y] = true
    for len(queue) != 0 {
       dist++ 
        for currLen := len(queue); currLen != 0; currLen-- {
            currNode, queue = queue[0], queue[1:]
            x, y := currNode[0], currNode[1]
            if x == numR -1 && y == numC -1 {
                return dist
            }
            for i := 0; i < 8; i++ {
                newX, newY := x + dirs[i][0], y + dirs[i][1]
                if (newX < 0 || newX >= numR || newY < 0 || newY >= numC) {
                    continue
                }
                if grid[newX][newY] == 1 {
                    continue
                }
                if visited[newX][newY] {
                    continue
                }
                queue = append(queue, [2]int{newX, newY})
                visited[newX][newY] = true
            }
        }
    }
    return -1
}