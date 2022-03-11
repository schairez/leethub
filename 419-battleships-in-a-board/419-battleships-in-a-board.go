func countBattleships(board [][]byte) int {
    const (
        ship = 'X'
        empty = '.'
    )
    numX := len(board)
    numY := len(board[0])
    visited := make([][]bool, numX)
    for x := range visited {
        visited[x] = make([]bool, numY)
    }
    dirs := [5]int{1,0,-1,0,1}
    var (
        res int
        dfs func(int, int) int
    )
    
    dfs = func(x, y int) int {
        if x < 0 || x >= numX || y < 0 || y >= numY{
            return 0
        }
        if board[x][y] == empty {
            return 0
        }
        if visited[x][y] {
            return 0
        }
        visited[x][y] = true
        for i := 1; i < len(dirs); i++ {
            deltaX, deltaY := dirs[i-1], dirs[i]
            dfs(x + deltaX, y + deltaY)
        }
        return  1
    }
    
    for x := 0; x < numX; x++ {
        for y := 0; y < numY; y++ {
            res += dfs(x, y)
        }
    }
    return res
}