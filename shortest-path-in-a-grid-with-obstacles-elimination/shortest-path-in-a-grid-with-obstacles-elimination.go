
/*
constraints
m == grid.length
n == grid[i].length
1 <= m, n <= 40
1 <= k <= m * n
grid[i][j] is either 0 or 1.
grid[0][0] == grid[m - 1][n - 1] == 0
*/

//goal:
// minimize distance s.t. 0 <= obstacles <= k
// BFS approach
// time: O(m*n)
// space: O(m*n)


type nodeData struct {
    x, y, cntObstacles int  
}

func shortestPath(grid [][]int, k int) int {
    numR := len(grid)
    numC := len(grid[0])
    //dstNodeR, dstNodeC := numR-1, numC-1 
    inArea := func(nextR, nextC int) bool {
        return nextR >=0 && numR > nextR && nextC >= 0 && numC > nextC
    }
    visited := make([][]int, numR)
    for row := range visited {
        visited[row] = make([]int, numC)
        for col := range visited[row] {
            visited[row][col] = k + 1
        }
    }
    visited[0][0] = 0
    //N, E, S, W
    xOffSet := [4]int{-1, 0, 1, 0}
    yOffSet := [4]int{0, 1, 0, -1}
    startNodeData := nodeData{0,0,0}
    //curR, curC, cntObstacles
    queue := []nodeData{startNodeData}
    numSteps := 0
    var cell nodeData
    for len(queue) != 0 {
        numNeighbors := len(queue)
        for nei := 0; nei < numNeighbors; nei++ {
            cell, queue = queue[0], queue[1:]
            if cell.cntObstacles ==  k+1 { continue }
            if cell.x == numR-1 && cell.y == numC-1 {
                return numSteps
            }
            x, y := cell.x, cell.y
            for dir :=0; dir < 4; dir++ {
                dx, dy := xOffSet[dir], yOffSet[dir]
                newX, newY := x + dx, y + dy
                if !inArea(newX, newY) { continue }
                //candidate for remaining obstacles; cntObstacles at cell
                candidate := cell.cntObstacles + grid[newX][newY]
                //if we already reached this node with less obstacles removed
                if visited[newX][newY] <= candidate { continue } 
                visited[newX][newY] = candidate
                queue = append(queue, nodeData{newX, newY, candidate})
            }
        }
        numSteps++
    }
    
    return -1
}