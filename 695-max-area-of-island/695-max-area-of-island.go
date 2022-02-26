func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}
func maxAreaOfIsland(grid [][]int) int {
    numR, numC := len(grid), len(grid[0])
    dirR := [4]int{-1, 0, 1, 0} 
    dirC := [4]int{0, 1, 0, -1}
    visited := make(map[[2]int]struct{})
    var res int
    for r := 0; r < numR; r++ {
        for c := 0; c < numC; c++ {
            if grid[r][c] == 1 {
                node := [2]int{r,c}
                if _, seen := visited[node]; !seen {
                    res = max(res, bfs(grid, visited, node, dirR, dirC))
                }
            }
        }
    }
    return res
}

func bfs(grid [][]int, visited map[[2]int]struct{}, node [2]int, dirR, dirC [4]int) int {
    numR, numC := len(grid), len(grid[0])
    visited[node] = struct{}{}
    var (
        queue [][2]int
        neiNode [2]int
        area int
        r, c int
    )
    queue = append(queue, node)
    for len(queue) != 0 {
        area++
        node, queue = queue[0], queue[1:]
        r, c = node[0], node[1]
        for i := 0; i < 4; i++ {
            nextR, nextC := r+dirR[i], c+dirC[i]
            if nextR < 0 || nextR >= numR || nextC < 0 || 
            nextC >= numC || grid[nextR][nextC] != 1 {
                continue
            }
            neiNode = [2]int{nextR, nextC}
            if _, seen := visited[neiNode]; !seen {
                queue = append(queue, neiNode)
                visited[neiNode] = struct{}{}
            }
        }
    }
    return area
}