func max(a, b int) int { if a >= b{ return a}; return b }
func min(a, b int) int { if a <= b { return a}; return b}

// todo: 
// k = max(numR, numC)?
// time: O(m*n*k) 
// space: O(m*n)
func maximumMinimumPath(grid [][]int) int {
    dirs := [5]int{1, 0,-1, 0, 1}
    
    numR, numC := len(grid), len(grid[0])
    dp := make([][]int, numR)
    for x := range dp {
        dp[x] = make([]int, numC)
    }
    inArea := func(x, y int) bool {
        return x >= 0 && x < numR && y >= 0 && y < numC
    }
    var (
        queue [][2]int
        currData [2]int
    )
    queue = append(queue, [2]int{0, 0})
    dp[0][0] = grid[0][0]
    for len(queue) != 0 {
        currData, queue = queue[0], queue[1:]
        x, y := currData[0], currData[1]
        for i := 1; i < 5; i++ {
            dX, dY := dirs[i-1], dirs[i]
            newX, newY := x + dX, y + dY
            if inArea(newX, newY)  {
                candScore := min(dp[x][y], grid[newX][newY])
                if candScore > dp[newX][newY] {
                    dp[newX][newY] = candScore
                    queue = append(queue, [2]int{newX, newY})
                }
            }
        } 
    }
    return dp[numR-1][numC-1]
    
}