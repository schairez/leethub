//multi-src BFS approach


func inArea(numR, numC int) func(int, int) bool {
    return func(i, j int) bool {
        if i < 0 || i >= numR || j < 0 || j >= numC {
            return false
        }
        return true
    }
}

    //seen := struct{}{}
    //visited := make(map[int][int]seen)

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
    //todo: collect num emptyRooms?
    numGates := 0
    for r := 0; r < numR; r++ {
        for c := 0; c < numC; c++ {
        if rooms[r][c] == gate {
            queue = append(queue, [2]int{r, c})
            numGates++
            }
        }
    }
    // TODO: if it's impossible to reach a gate fill with inf
    //if numGates == 0 {
        //return rooms or fill with inf?
    //} 
    var r, c int
    var dist int
    for len(queue) != 0 {
        r, c = queue[0][0], queue[0][1]
        dist = rooms[r][c]
        queue = queue[1:]
        //visited[r][c] = seen
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