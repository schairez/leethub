func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    ufPar := make([]int, n)
    ufSize := make([]int, n)
    for x := range ufSize {
        ufSize[x] = 1
        ufPar[x] = x
    }
    numForests := n
    findRoot := func(x int) int {
        for x != ufPar[x] {
            ufPar[x] = ufPar[ufPar[x]]
            x = ufPar[x]
        }
        return x
    }
    union := func(w, v int) bool {
        rootW, rootV := findRoot(w), findRoot(v)
        if rootW == rootV {
            return false
        }
        wSze, vSze := ufSize[rootW], ufSize[rootV]
        if wSze <= vSze {
            ufSize[vSze] += wSze
            ufPar[rootW] = rootV
        } else {
            ufSize[wSze] += vSze
            ufPar[rootV] = rootW
        }
        return true
    }
    
    for x := 0; x < n; x++ {
        for y := 0; y < n; y++ {
            if isConnected[x][y] == 1 && union(x, y) {
                numForests--
            }
        }
    }
    
    
    return numForests
    
}