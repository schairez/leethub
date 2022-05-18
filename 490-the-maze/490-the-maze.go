func hasPath(maze [][]int, start []int, destination []int) bool {
    // condition: start.length() == dest.length() == 2
    numR, numC := len(maze), len(maze[0])
    inArea := func(x, y int) bool {
        return x >= 0 && y >= 0 && x < numR && y < numC 
    }
    
    visited := make([][]bool, numR)
    for x := range visited {
        visited[x] = make([]bool, numC)
    }
    dirs := [5]int{1,0,-1,0,1}
    var (
        x, y  int
        node  [2]int
        queue [][2]int
    )
    x, y = start[0], start[1]
    visited[x][y] = true
    queue = append(queue, [2]int{x, y})
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            x, y = node[0], node[1]
            if x == destination[0] && y == destination[1] {
                return true
            }
            for i := 1; i < 5; i++ {
                newX, newY := x, y
                for inArea(newX + dirs[i-1], newY + dirs[i]) &&
                maze[newX + dirs[i-1]][newY + dirs[i]] == 0 {
                    newX += dirs[i-1]
                    newY += dirs[i]
                }
                if !visited[newX][newY] {
                    queue = append(queue, [2]int{newX, newY})
                    visited[newX][newY] = true
                }
            }
        }
    }
    return false    
}