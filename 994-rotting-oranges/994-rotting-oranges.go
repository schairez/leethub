// time: O(m * n)
// space: O(m * n)

const (
    empty = iota
    fresh
    rotten
)

func orangesRotting(grid [][]int) int {
    numR, numC := len(grid), len(grid[0])
    var queue [][2]int
    var numFresh int
    for r := 0; r < numR; r++ {
        for c := 0; c < numC; c++ {
            if grid[r][c] == rotten {
                queue = append(queue, [2]int{r,c})
            } else if grid[r][c] == fresh {
                numFresh++
            } 
        }
    }
    // (r-1, c)
    dirR := [4]int{-1, 0, 1, 0}
    dirC := [4]int{0, 1, 0, -1}
    var (
        node [2]int
        r, c int
        lvl int
        rottenSoFar int
    )
    for len(queue) != 0 && rottenSoFar != numFresh {
        for currLen := len(queue); currLen != 0; currLen-- {
            node, queue = queue[0], queue[1:]
            r, c = node[0], node[1]
            for i := 0; i < 4; i++ {
                nextR, nextC := r + dirR[i], c + dirC[i]
                if nextR >= 0 && nextR < numR && 
                nextC >= 0 && nextC < numC &&
                grid[nextR][nextC] == fresh {
                    grid[nextR][nextC] = rotten
                    queue = append(queue, [2]int{nextR, nextC} )
                    rottenSoFar++
                }
            }
        }
        lvl++
    }
    
    if rottenSoFar != numFresh {
        return -1
    }
    return lvl
}