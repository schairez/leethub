
// 562. Longest Line of Consecutive One in Matrix
// time: O(m*n*4) ≈ O(m*n)
// space: O(m*n*4) ≈ O(m*n)

func max(a, b int) int { if a >= b { return a}; return b}

func longestLine(mat [][]int) int {
    numR, numC := len(mat), len(mat[0])
    if numR == 1 && numC == 1 {
        return mat[numR-1][numC-1]
    }
    dp := make([][][4]int, numR)
    for x := range dp {
        dp[x] = make([][4]int, numC)
    }
    // The line could be horizontal, vertical, diagonal, or anti-diagonal.
    dirs := [4][2]int{{0, -1}, {-1, 0}, {-1, -1}, {-1, 1}} 
    res := 0
    for x := 0; x < numR; x++ {
        for y := 0; y < numC; y++ {
            if mat[x][y] == 0 {
                continue
            }
            for idx := range dirs {
                dp[x][y][idx] = 1
                prevX, prevY := x + dirs[idx][0], y + dirs[idx][1]
                if prevX < 0 || prevX >= numR || prevY < 0 || prevY >= numC {
                    continue
                }
                dp[x][y][idx] += dp[prevX][prevY][idx]
                res = max(res, dp[x][y][idx])
            }
        }
    }
    return res
}