func shortestDistance(maze [][]int, start []int, destination []int) int {
    const maxInt32 = 1 << 31 - 1
    dirs := [5]int{1, 0, -1, 0, 1}
    numR, numC := len(maze), len(maze[0])
    distGrid := make([][]int, numR)
    for x := range distGrid {
        distGrid[x] = make([]int, numC)
        for y := range distGrid[0] {
            distGrid[x][y] = maxInt32
        }
    }
    inArea := func(x, y int) bool {
        return x >= 0 && x < numR && y >= 0 && y < numC
    }
    
    var (
        queue [][2]int
        node  [2]int 
        x, y  int
    )
    x, y = start[0], start[1]
    distGrid[x][y] = 0
    queue = append(queue, [2]int{x, y})
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            x, y = node[0], node[1]
            for i := 1; i < 5; i++ {
                newX, newY := x, y
                incrDist := 0
                for inArea(newX + dirs[i-1], newY + dirs[i]) &&
                maze[newX + dirs[i-1]][newY + dirs[i]] == 0 {
                    newX += dirs[i-1]
                    newY += dirs[i]
                    incrDist++
                }
                if candDist := distGrid[x][y] + incrDist; distGrid[newX][newY] > candDist  {
                    distGrid[newX][newY] = candDist
                    queue = append(queue, [2]int{newX, newY})
                }
            }
        }
    }
    if res := distGrid[destination[0]][destination[1]]; res != maxInt32 {
        return res
    }
    return -1
}