const (
    emptyLand = iota
    building
    obstacle
)

const (
    maxInt32 = 1 << 31 -1
)

func shortestDistance(grid [][]int) int {
    numR, numC := len(grid), len(grid[0])
    
    total := 0
    for r := 0; r < numR; r++ {
        for c := 0; c < numC; c++ {
            if grid[r][c] == building {
                total++
            }
        }
    }
    var minDist int = maxInt32
    for r := 0; r < numR; r++ {
        for c := 0; c < numC; c++ {
            if grid[r][c] == emptyLand {
                if candDist := bfs(grid, total, r, c); candDist < minDist {
                    minDist = candDist
                }
            } 
        }
    }
    if minDist == maxInt32 {
        return -1
    }
    return minDist
}

func inArea(numR, numC, r, c int) bool {
    if r < 0 || r >= numR || c < 0 || c >= numC {
        return false
    }
    return true
}

func bfs(grid [][]int, total, startR, startC int) int {
    dX := [4]int{0,0,1,-1}
    dY := [4]int{1,-1,0,0}
    numR, numC := len(grid), len(grid[0])
    totalDist, dist := 0, 0
    numBuildings := 0
    visited := make([][]bool, numR)
    for r := range visited {
        visited[r] = make([]bool, numC)
    }
    queue := make([][2]int, 0)
    queue = append(queue, [2]int{startR,startC})
    var node [2]int
    var r, c int
    for len(queue) != 0 && numBuildings != total {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            r, c = node[0], node[1]
            if grid[r][c] == building {
                numBuildings++
                totalDist += dist
                continue
            } 
            for idx := 0; idx < 4; idx++ {
                nextR, nextC := r + dX[idx], c + dY[idx]
                if inArea(numR, numC, nextR, nextC) && !visited[nextR][nextC] &&
                grid[nextR][nextC] != obstacle {
                    queue = append(queue, [2]int{nextR, nextC})
                    visited[nextR][nextC] = true
                }
            }
        }
        dist++
    }
    if numBuildings != total {
        return maxInt32
    }
    return totalDist
} 
