
//multi-src BFS approach
// time: O(m*n) to find src gate nodes
// space: O(m*n) since we can have a case where we insert a high number
// of gates at once, and depends on growth of our rooms input

func inArea(numR, numC int) func(int, int) bool {
    return func(i, j int) bool {
        if i < 0 || i >= numR || j < 0 || j >= numC {
            return false
        }
        return true
    }
}


func wallsAndGates(rooms [][]int)  {
    numR, numC := len(rooms), len(rooms[0])
    inAreaFn := inArea(numR, numC)
    // constraint: 1 <= m, n <= 250
    // queue up all the src nodes
    // here srcNodes are (i,j) that represents a gate
    const (
        gate = 0
        inf = 1 << 31 - 1
    )
    var queue [][2]int
    for r := 0; r < numR; r++ {
        for c := 0; c < numC; c++ {
        if rooms[r][c] == gate {
            queue = append(queue, [2]int{r, c})
            }
        }
    }
    var r, c int
    var dist int
    for len(queue) != 0 {
        r, c = queue[0][0], queue[0][1]
        dist = rooms[r][c]
        queue = queue[1:]
        // N
        if inAreaFn(r-1, c) && rooms[r-1][c] == inf {
            rooms[r-1][c] = dist+1
            queue = append(queue, [2]int{r-1, c})
        } 
        // E
        if inAreaFn(r, c+1) && rooms[r][c+1] == inf {
            rooms[r][c+1] = dist+1
            queue = append(queue, [2]int{r, c+1})
        } 
        // S
        if inAreaFn(r+1, c) && rooms[r+1][c] == inf {
            rooms[r+1][c] = dist+1
            queue = append(queue, [2]int{r+1, c})
        } 
        // W
        if inAreaFn(r, c-1) && rooms[r][c-1] == inf {
            rooms[r][c-1] = dist+1
            queue = append(queue, [2]int{r, c-1})
        } 
    }
}