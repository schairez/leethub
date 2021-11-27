//time: O(N); we iterate over all cells in matrix
//space: O(N)


/*
All the adjacent cells of the path are 8-directionally connected
(i.e., they are different and they share an edge or a corner).

clear path is path from
(0,0) to (n - 1, n - 1)

1 <= n <= 100
n == grid.length
n == grid[i].length
grid[i][j] is 0 or 1

*/


func shortestPathBinaryMatrix(grid [][]int) int {
    n := len(grid)
    if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
        return -1
    } 
    
    dirs := [8][2]int{{-1,0},//N
                      {1,0},//S
                      {0,1},//E
                      {0,-1},//W
                      {-1,-1},//NW
                      {-1,1},//NE
                      {1,-1},//SW
                      {1,1},//SE
                     }
    visited := make([][]bool, n)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    pathDistance := 1
    queue := [][2]int{[2]int{0,0}}
    for len(queue) != 0 {
        nodesAtCurrDist := len(queue)
        for i :=0; i < nodesAtCurrDist; i++ {
            var pollNode [2]int
            pollNode, queue = queue[0], queue[1:]
            currX, currY := pollNode[0], pollNode[1]
            if currX == n-1 && currY == n-1 {
                return pathDistance
            }
            for _, dir := range dirs {
                dx, dy := dir[0], dir[1]
                newX, newY := currX +dx, currY+dy
                if !(newX >= 0 && newX < n && newY >= 0 && newY < n) || 
                visited[newX][newY] || grid[newX][newY] != 0 {
                    continue
                } 
                queue = append(queue, [2]int{newX,newY})
                visited[newX][newY] = true
            }
        } 
        pathDistance++
    }
    return -1
}









