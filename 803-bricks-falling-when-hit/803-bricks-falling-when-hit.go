// space: O(numR*numC)
// time: O(numR*numC*numHits* a(n))
//by inverseAckerman find ops are O(a(n)) where a(n) < 5


const (
    emptySpace = iota
    brick
)

func hitBricks(grid [][]int, hits [][]int) []int {
    numX, numY := len(grid), len(grid[0])
    dirX := [4]int{-1, 0, 1, 0}
    dirY := [4]int{0, -1, 0, 1}
    
    inArea := func(x, y int) bool {
        return x >= 0 && x < numX && y >= 0 && y < numY
    }
    getNodeID := func(x, y int) int {
        return x * numY + y
    } 
    
    ufSize := make([]int, numX * numY + 1)
    ufParent := make([]int, numX * numY + 1)
    for id := range ufSize {
        ufSize[id] = 1
        ufParent[id] = id
    } 
    findRoot := func(v1 int) int {
        for v1 != ufParent[v1] {
            ufParent[v1] = ufParent[ufParent[v1]]
            v1 = ufParent[v1]
        }
        return v1
    }
    union := func(v1, v2 int) {
        v1Root, v2Root := findRoot(v1), findRoot(v2)
        if v1Root == v2Root {
            return
        }
        if ufSize[v1Root] <= ufSize[v2Root] {
            ufParent[v1Root] = v2Root
            ufSize[v2Root] += ufSize[v1Root]
        } else {
            ufParent[v2Root] = v1Root
            ufSize[v1Root] += ufSize[v2Root]
        }
    }
    
    gridCopy := make([][]int, numX)
    for x := range gridCopy {
        //gridCopy[x] = make([]int, numY)
        gridCopy[x] = append([]int(nil), grid[x]...)
    }
    for _, hit := range hits {
        x, y := hit[0], hit[1]
        gridCopy[x][y] = emptySpace 
    }
    rootID := numX * numY
    for x := 0; x < numX; x++ {
        for y := 0; y < numY; y++ {
            if gridCopy[x][y] == emptySpace {
                continue
            }
            if x == 0 {
                union(getNodeID(x,y), rootID)
            }
            // from above
            if x > 0 && gridCopy[x-1][y] == brick {
                union(getNodeID(x,y), getNodeID(x-1, y))
            }
            // from left
            if y > 0 && gridCopy[x][y-1] == brick {
                union(getNodeID(x,y), getNodeID(x, y-1))
            }
        }
    }
    numHits := len(hits)
    res := make([]int, numHits)
    for idx := numHits-1; idx >= 0; idx-- {
        x, y := hits[idx][0], hits[idx][1]
        if grid[x][y] == 0 {
            continue
        }
        prevSize := ufSize[findRoot(rootID)] 
        if x == 0 {
            union(getNodeID(x, y), rootID)
        }
        
        for dir := 0; dir < 4; dir++ {
            nextX, nextY:= x + dirX[dir], y + dirY[dir]
            if !inArea(nextX, nextY) || 
            gridCopy[nextX][nextY] == emptySpace {
                continue
            }
            union(getNodeID(x,y), getNodeID(nextX, nextY))
        }
        currSize := ufSize[findRoot(rootID)] 
        if numFell := currSize - prevSize -1; numFell > 0 {
            res[idx] = numFell
        }
        gridCopy[x][y] = brick  
    }
    return res
    
}
