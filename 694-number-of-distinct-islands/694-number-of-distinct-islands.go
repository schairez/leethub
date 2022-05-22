
// 694. Number of Distinct Islands
// time: O(m*n)
// space: O(m*n)

func numDistinctIslands(grid [][]int) int {
    numR, numC := len(grid), len(grid[0])
    inArea := func(x, y int) bool {
        return x >= 0 && x < numR && y >= 0 && y < numC
    }
    visited := make([][]bool, numR)
    for x := range visited {
        visited[x] = make([]bool, numC)
    }
    distinctIslands := make(map[string]struct{})
    var dfs func(int, int, byte, *strings.Builder)
    dfs = func(x, y int, dir byte, sb *strings.Builder) {
        if !inArea(x, y) || grid[x][y] != 1 || visited[x][y] {
            return
        }
        visited[x][y] = true
        sb.WriteByte(dir)
        dfs(x - 1, y, 'N', sb)
        dfs(x + 1, y, 'S', sb)
        dfs(x, y + 1, 'E', sb)
        dfs(x, y - 1, 'W', sb)
        sb.WriteByte('B')
    }
    
    for x := 0; x < numR; x++ {
        for y := 0; y < numC; y++ {
            if grid[x][y] == 1 && !visited[x][y] {
                sb := &strings.Builder{}
                dfs(x, y, '$', sb)
                cand := sb.String()
                if _, exists := distinctIslands[cand]; !exists {
                    distinctIslands[cand] = struct{}{}
                }
            }
        }
    }
    
    return len(distinctIslands)
}