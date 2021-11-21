//time: O(m*n)
//space: O(1)

func islandPerimeter(grid [][]int) int {
    res := 0
    numRows := len(grid)
    numCols := len(grid[0])
    for r:=0; r< numRows; r++ {
        for c:=0; c < numCols; c++ {
            if grid[r][c]==1 {
                res +=4
                if c > 0 && grid[r][c] == grid[r][c-1] {
                    res -= 2
                }
                if r > 0 && grid[r][c] == grid[r-1][c] {
                    res -= 2
                }
            }
        }
    }
    return res
}    











/*
//nonworking below; failed first attempt

func islandPerimeter(grid [][]int) int {
    numRows := len(grid)
    numCols := len(grid[0])
    visited := make([][]bool, numRows)
    for i := range visited {
        visited[i] = make([]bool, numCols)
    }
    inArea := func(x, y int) bool {
        if (x >= 0 && x <= numRows-1) && (y >= 0 && y <= numCols-1) {
            return true
        } 
        return false
    }
    bfs := func(startCell [2]int) int {
        //numConnComp := 0
        dirs := [4][2]int{{1,0},  //N
                          {0,1},  //E
                          {-1,0}, //S
                          {0,-1}, //W
                         }
        //perim := 0
        numConnComp := 0
        neighbors := 0
        q := [][2]int{}
        q = append(q, startCell)
        for len(q) > 0 {
            node := q[0]
            q = q[1:]
            numConnComp += 1
            x, y := node[0], node[1]
            for _, dir := range dirs {
                dx, dy := dir[0], dir[1]
                newX, newY := x + dx, y + dy
                if !inArea(newX, newY) {
                    continue
                }
                if !visited[newX][newY] && grid[newX][newY] == 1 {
                    q = append(q, [2]int{newX, newY})
                    visited[newX][newY] = true
                    neighbors++
                    //numConnComp++
                }  
            }
            //perim += (4 - numConnComp)
            visited[x][y] = true
        }
        //return perim
        fmt.Println(numConnComp)
        fmt.Println(neighbors)
        
        return numConnComp*4 - 2*neighbors
        //return (numConnComp * 4) - 2*(numConnComp-1)
    }
    
    connComp := [2]int{}
    var res int
    for r:=0; r < numRows; r++ {
        for c:=0; c < numCols; c++ {
            if grid[r][c] == 1 {
                connComp = [2]int{r, c}
                res := bfs(connComp)
                return res
            }
        }
    }
    return res
    
}
*/