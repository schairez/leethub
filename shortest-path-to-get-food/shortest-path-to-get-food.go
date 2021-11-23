//time:O(m*n)
//space:O(m*n)

func getFood(grid [][]byte) int {
    numR := len(grid)
    numC := len(grid[0])
    //The grid contains exactly one '*'.
    findStartLoc := func() (int,int) {
        var startR, startC int
        notFoundLoop:
        for i:=0; i<numR; i++ {
            for j:=0; j<numC; j++ {
                if grid[i][j] == byte('*') {
                    startR, startC = i, j
                    break notFoundLoop
                }
            }
        }  
        return startR, startC
    }
    startR, startC := findStartLoc() 
    
    
    visited := make([][]bool, numR)
    for i := range visited {
        visited[i] = make([]bool, numC)
    }
    visited[startR][startC] = true
    dirRowOps := [4]int{0, 1, 0, -1}
    dirColOps := [4]int{1, 0, -1, 0}
    queue := [][2]int{}
    queue = append(queue, [2]int{startR, startC})
    level := 0
    
    for len(queue) > 0 {
        curLen := len(queue)
        for i:=0; i < curLen; i++ {
            poll := queue[0]
            queue = queue[1:]
            x, y := poll[0], poll[1]
            //if grid[x][y] == byte('#') { return level }
            //travel to adj cell
            for cell:=0; cell < 4; cell++ { 
                newX, newY := x + dirRowOps[cell], y + dirColOps[cell]
                //check if (x,y) inArea
                if !(newX >= 0 && newX <= numR-1 && newY >=0 && newY <= numC-1) {
                    continue
                } 
                //already visited or an obstacle
                if visited[newX][newY] || grid[newX][newY] == byte('X') { continue }
                if grid[newX][newY] == byte('#') {
                    return level+1
                }
                queue = append(queue, [2]int{newX, newY})
                visited[newX][newY] = true
            }
        }
        level++
    } 
    return -1
}


/*
                   loc  obst   free    food
grid[row][col] is '*', 'X',   'O', or  '#'.

grid = 
[["X","X","X","X","X","X"],
["X","*","O","O","O","X"],
["X","O","O","#","O","X"],
["X","X","X","X","X","X"]]

*/
