// 361. Bomb Enemy
// bfs approach
// time: O(m*n + (m+n))
// space: O(1)

func max(a, b int) int {if a >=b {return  a}; return b}

func maxKilledEnemies(grid [][]byte) int {
    const (
        wall  byte = 'W'
        enemy byte = 'E'
        empty byte = '0'
    )
    dirs := [5]int{1,0,-1,0,1}
    numR, numC := len(grid), len(grid[0])
    isValid := func(x, y int) bool {
        return x >= 0 && y >= 0 && y < numC && x < numR
    }
    // O(m+n)
    bfs := func(x, y int) int {
        numEnemies := 0
        for i := 1; i < 5; i++ {
            newX, newY := x, y
            for isValid(newX, newY) && grid[newX][newY] != wall {
                if grid[newX][newY] == enemy {
                    numEnemies++
                }
                newX, newY = newX + dirs[i-1], newY + dirs[i]
            }
        }
        return numEnemies
    }
    res := 0
    // O(m*n)
    for r := 0; r < numR; r++ {
        for c := 0; c < numC; c++ {
            if grid[r][c] == empty {
                res = max(res, bfs(r,c))
            }
        }
    }
    return res
}