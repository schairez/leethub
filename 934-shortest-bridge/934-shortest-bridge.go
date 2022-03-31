// 934. Shortest Bridge
// time: O(m*n)
// space: O(m*n)

// There are exactly two islands in grid.

func shortestBridge(grid [][]int) int {
    dirs := [5]int{1, 0, -1, 0, 1}
    numR, numC := len(grid), len(grid[0])
    //visited := make(map[[2]int]struct{})
    visited := make([][]bool, numR)
    for x := range visited {
        visited[x] = make([]bool, numC)
    }
    var (
        queue [][2]int
        node  [2]int
        lvl   int
    )
    isOneIslandFound := false
    for x := 0; x < numR; x++ {
        if isOneIslandFound {
            break
        }
        for y := 0; y < numC; y++ {
            if grid[x][y] == 1 && !visited[x][y] {
                visited[x][y] = true
                var tmpQueue [][2]int
                node = [2]int{x,y}
                tmpQueue = append(tmpQueue, node)
                for len(tmpQueue) != 0 {
                    node, tmpQueue = tmpQueue[0], tmpQueue[1:]
                    queue = append(queue, node)
                    x, y := node[0], node[1]
                    for i := 1; i < 5; i++ {
                        newX, newY := x + dirs[i-1], y + dirs[i]
                        if newX < 0 || newX >= numR || newY < 0 || newY >= numR {
                            continue
                        }
                        if grid[newX][newY] == 0 {
                            continue
                        }
                        if visited[newX][newY] {
                            continue
                        }
                        tmpQueue = append(tmpQueue, [2]int{newX, newY})
                        visited[newX][newY] = true
                    }
                }
                isOneIslandFound = true
                break
            }
        }
    }
    
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            x, y := node[0], node[1]
            for i := 1; i < 5; i++ {
                newX, newY := x + dirs[i-1], y + dirs[i]
                if newX < 0 || newX >= numR || newY < 0 || newY >= numR {
                    continue
                }
                if visited[newX][newY] {
                    continue
                }
                if grid[newX][newY] == 1 {
                    return lvl
                }
                queue = append(queue, [2]int{newX, newY})
                visited[newX][newY] = true
            }
        }
        lvl++
    }
    
    return -1
    
}