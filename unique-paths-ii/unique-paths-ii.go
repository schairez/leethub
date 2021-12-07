//top-down memoized version
//time: O(m*n)
//space: O(m*n)

func uniquePathsWithObstaclesMemo(obstacleGrid [][]int) int {
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

//bottom up DP 2D
//time: O(m*n)
//space: O(m*n)
func uniquePathsWithObstacles2D(obstacleGrid [][]int) int {
    numR, numC := len(obstacleGrid), len(obstacleGrid[0])
    if obstacleGrid[0][0] == 1 { return 0 }
    dp := make([][]int, numR)
    for row := range dp {
        dp[row] = make([]int, numC)
    }
    for row := range dp {
        if obstacleGrid[row][0] == 1 {
            break
        }
        dp[row][0] = 1
    }
    for col := 0; col < numC; col++ {
        if obstacleGrid[0][col] == 1 {
            break
        }
        dp[0][col] = 1
    }
    for row := 1; row < numR; row++ {
        for col := 1; col < numC; col++ {
            if obstacleGrid[row][col] == 1 {
                dp[row][col] = 0
                continue
            }
            dp[row][col] = dp[row-1][col] + dp[row][col-1]
        }
    }
    return dp[numR-1][numC-1]
}
//bottom-up 1D
//time: O(m*n)
//space: O(n)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    numR, numC := len(obstacleGrid), len(obstacleGrid[0])
    if obstacleGrid[0][0] == 1 { return 0 }
    dp := make([]int, numC)
    //iter through row = 0 col vals
    for col := 0; col < numC; col++ {
        if obstacleGrid[0][col] == 1 {
            break
        }
        dp[col] = 1
    }
    for row := 1; row < numR; row++ {
        //val at prev row
        prev := dp[0]
        if obstacleGrid[row][0] == 1 {
            dp[0] = 0
        } else {
            dp[0] = prev
        }
        for col := 1; col < numC; col++ {
            prev = dp[col]
            if obstacleGrid[row][col] == 1 {
                dp[col] = 0
            } else {
                dp[col] = prev + dp[col-1]
            }
        }
    }
    return dp[numC-1]
}