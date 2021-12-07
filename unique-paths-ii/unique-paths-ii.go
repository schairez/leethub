//top-down memoized version
//time: O(m*n)
//space: O(m*n)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    numR, numC := len(obstacleGrid), len(obstacleGrid[0])
    if obstacleGrid[0][0] == 1 { return 0 }
    memo := make([][]int, numR)
    for row := range memo {
        memo[row] = make([]int, numC)
    }
    var dfs func(int, int) int
    dfs = func(cellR, cellC int) int {
        if cellR < 0 || cellC < 0 || obstacleGrid[cellR][cellC] == 1 {
            return 0
        }
        if cellR == 0 && cellC == 0 {
            return 1
        }
        if numPaths := memo[cellR][cellC]; numPaths != 0 {
            return numPaths
        } 
        //recursive from left and from right
        memo[cellR][cellC] = dfs(cellR-1, cellC) + dfs(cellR, cellC-1)
        return memo[cellR][cellC]
    }
    
    return dfs(numR-1, numC-1)
}
