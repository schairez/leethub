
const (
    emptyLand = iota
    house
    obstacle
)

// 317. Shortest Distance from All Buildings
// BFS from buildings to empty lands approach
// time: O(m^2 * n^2)
// space: O(m*n)

func shortestDistance(grid [][]int) int {
    numX, numY := len(grid), len(grid[0])
    // condition: 1 <= m, n <= 50
    // enqueue all the buildings
    var (
        houseCoord [2]int
        houseQueue [][2]int
    )
    for x := 0; x < numX; x++ {
        for y := 0; y < numY; y++ {
            if grid[x][y] == house {
                houseQueue = append(houseQueue, [2]int{x, y})
            }
        }
    }
    // constraint: There will be at least one building in the grid.
    totalHouses := len(houseQueue)
    
    distanceGrid := make([][]int, numX)
    numReachGrid := make([][]int, numX)
    for x := range distanceGrid {
        distanceGrid[x] = make([]int, numY)
        numReachGrid[x] = make([]int, numY)
    }
    numReachable := 1
    for len(houseQueue) != 0 {
        houseCoord, houseQueue = houseQueue[0], houseQueue[1:]
        bfs(grid, distanceGrid, numReachGrid, houseCoord, numReachable)
        numReachable++
    }
    const maxInt32 = 1 << 31 - 1
    minDist := maxInt32
    for x := 0; x < numX; x++ {
        for y := 0; y < numY; y++ {
            if numReachGrid[x][y] == totalHouses &&
            distanceGrid[x][y] < minDist {
                minDist = distanceGrid[x][y]
            }
        }
    }
    if minDist == maxInt32 {
        return -1
    }
    return minDist
}

func bfs(grid [][]int, distanceGrid [][]int,
         numReachGrid[][]int, houseCoord [2]int, numReachable int ) {
    
    numX, numY := len(grid), len(grid[0])
    dirs := [5]int{1, 0, -1, 0, 1}
    var (
        node [2]int
        queue [][2]int
    )
    dist := 0
    queue = append(queue, houseCoord)
    for len(queue) != 0 {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            x, y := node[0], node[1]
            distanceGrid[x][y] += dist
            for i := 1; i < 5; i++ {
                newX := x + dirs[i-1]
                newY := y + dirs[i]
                if newX < 0 || newX >= numX || newY < 0 || newY >= numY {
                    continue
                }
                if grid[newX][newY] == house || grid[newX][newY] == obstacle {
                    continue
                }
                if numReachGrid[newX][newY] == numReachable - 1 {
                    queue = append(queue, [2]int{newX, newY})
                    numReachGrid[newX][newY] = numReachable
                }
            }
        }
        dist++
    }
}

