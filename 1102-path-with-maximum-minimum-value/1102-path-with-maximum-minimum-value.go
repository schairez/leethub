
func max(a, b int) int { if a >= b{ return a}; return b }
func min(a, b int) int { if a <= b { return a}; return b}

func maximumMinimumPath(grid [][]int) int {
    numR, numC := len(grid), len(grid[0])
    inArea := func(x, y int) bool {
        return x >= 0 && x < numR && y >= 0 && y < numC
    }
    dirs := [5]int{1, 0, -1, 0, 1}
    pq := newPQ(func(a, b int) bool {
        return a > b 
    })
    heap.Init(pq)
    visited := make([][]bool, numR)
    for x := range visited {
        visited[x] = make([]bool, numC)
    }
    heap.Push(pq, pathData{0, 0, grid[0][0]})
    visited[0][0] = true
    
    res := grid[0][0]
    
    for pq.Len() != 0 {
        curr := heap.Pop(pq).(pathData)
        res = min(res, curr.score)
        if curr.x == numR-1 && curr.y == numC-1 {
            break
        }
        for i := 1; i < 5; i++ {
            newX, newY := curr.x + dirs[i-1], curr.y + dirs[i]
            if inArea(newX, newY) && !visited[newX][newY] {
                data := pathData{newX, newY, min(curr.score, grid[newX][newY])}
                heap.Push(pq, data)
                visited[newX][newY] = true
            }
        }
    }
    
    return res
}



type pathData struct {
    x, y, score int
}

type PriorityQueue struct {
    data []pathData
    lessFn func(i, j int) bool
}
func newPQ(lessFn func(i, j int) bool) *PriorityQueue {
    return &PriorityQueue {
        lessFn: lessFn,
    }
}

func (pq *PriorityQueue) Len() int {
    return len(pq.data)
}
func (pq *PriorityQueue) Less(i, j int) bool {
    return pq.lessFn(pq.data[i].score, pq.data[j].score)
}
func (pq *PriorityQueue) Swap(i, j int) {
    pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}
func (pq *PriorityQueue) Push(v interface{}) {
    pq.data = append(pq.data, v.(pathData))
}
func (pq *PriorityQueue) Pop() interface{} {
    n := pq.Len()
    v := pq.data[n-1]
    pq.data = pq.data[:n-1]
    return v
}




// BFS + DP approach
// k = max(numR, numC)
// time: O(m*n*k) 
// space: O(m*n)
func maximumMinimumPathDP(grid [][]int) int {
    dirs := [5]int{1, 0,-1, 0, 1}
    
    numR, numC := len(grid), len(grid[0])
    dp := make([][]int, numR)
    for x := range dp {
        dp[x] = make([]int, numC)
    }
    inArea := func(x, y int) bool {
        return x >= 0 && x < numR && y >= 0 && y < numC
    }
    var (
        queue [][2]int
        currData [2]int
    )
    queue = append(queue, [2]int{0, 0})
    dp[0][0] = grid[0][0]
    for len(queue) != 0 {
        currData, queue = queue[0], queue[1:]
        x, y := currData[0], currData[1]
        for i := 1; i < 5; i++ {
            dX, dY := dirs[i-1], dirs[i]
            newX, newY := x + dX, y + dY
            if inArea(newX, newY)  {
                candScore := min(dp[x][y], grid[newX][newY])
                if candScore > dp[newX][newY] {
                    dp[newX][newY] = candScore
                    queue = append(queue, [2]int{newX, newY})
                }
            }
        } 
    }
    return dp[numR-1][numC-1]
    
}