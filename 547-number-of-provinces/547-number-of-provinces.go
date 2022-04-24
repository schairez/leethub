// 547. Number of Provinces
// time: O(n*n)
// space: O(n)

func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    numProvinces := n
    parent := make([]int, n)
    sze := make([]int, n)
    for x := range sze {
        sze[x] = 1
        parent[x] = x
    }
    // path compression
    // inv ackerman O(a(n)) <= 4
    findRoot := func(x int) int {
        for x != parent[x] {
            parent[x] = parent[parent[x]]
            x = parent[x]
        }
        return x
    }
    // union by size with path compression
    union := func(w, v int) bool {
        wRoot := findRoot(w)
        vRoot := findRoot(v)
        if wRoot == vRoot {
            return false
        }
        // small into big
        if sze[wRoot] <= sze[vRoot] {
            sze[vRoot] += sze[wRoot]
            parent[wRoot] = vRoot
        } else { // wRoot > vRoot
            sze[wRoot] += sze[vRoot]
            parent[vRoot] = wRoot
        }
        return true
    }
    for x := 0; x < n; x++ {
        for y := 0; y < n; y++ {
            if isConnected[x][y] == 1 && union(x, y) {
                numProvinces--
            }
        }
    }
    return numProvinces
    
}