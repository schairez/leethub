

func swimInWater(grid [][]int) int {
    dirs := [4][2]int{{1,0}, //north
                      {-1,0}, //south
                      {0,1}, //east
                      {0,-1}, //west
                     }
    //1 <= n <= 50
    //0 <= grid[i][j] < n2
    //no issue of int overflow with such small size
    n := len(grid)
    //for binSearch
    lo, hi := 0, n*n
    for lo < hi {
        mid := (lo + hi) >> 1 
        if checkPath(grid, dirs,mid, n) {
            hi = mid
        } else {
            lo = mid + 1
        }
    }
    return hi
}

func max(a, b int) int { if a >= b { return a}; return b }


func checkPath(grid [][]int, dirs[4][2]int, time, n int) bool {
    inGridArea := func(n, x, y int) bool {
        if (x >=0 && x <= n-1) && (y >=0 && y <= n-1) {
            return true
        }  
        return false
    }
    canMove := func(time, srcX, srcY, dstX, dstY int ) bool {
        //You can swim from a square to another 4-directionally
        //adjacent square iff the elevation of both squares individually are at most t.
        return time >= max(grid[srcX][srcY], grid[dstX][dstY]) 
    }
    visited := make([][]bool, n)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    visited[0][0] = true
    queue := [][2]int{}
    queue = append(queue, [2]int{0,0})
    for len(queue) > 0 {
        pollNode := queue[0]
        queue = queue[1:]
        srcX, srcY := pollNode[0], pollNode[1]
        if srcX == n-1 && srcY == n-1 { return true }
        
        for _, dir := range dirs {
            dx, dy := dir[0], dir[1]
            dstX, dstY := srcX + dx, srcY + dy
            if !inGridArea(n, dstX, dstY) { continue }
            if !visited[dstX][dstY] && canMove(time, srcX, srcY, dstX, dstY) {
                visited[dstX][dstY] = true
                queue = append(queue, [2]int{dstX, dstY} )
            } 
        }
    }
    return false
}



