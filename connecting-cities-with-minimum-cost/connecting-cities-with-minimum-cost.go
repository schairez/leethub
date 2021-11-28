/*

N cities and E edges/connections
time: O(E*logE); for sorting ElogE; for uf union and find operations
the find takes logN time, but since the runtime of this operation never exceeds 5 for 10^600 ops, we can see this as virtually O(1)
therefore O(ElogE) is the runtime
space: O(2N) ~ O(n)

*/

import "sort"

type ByWeight [][]int

func (w ByWeight) Len() int {
    return len(w)
}
func (w ByWeight) Swap(i, j int) {
    w[i], w[j] = w[j], w[i]
}
func (w ByWeight) Less(i, j int) bool {
    if w[i][2] < w[j][2] {
        return true
    }
    return false
}

type UF struct {
    Par []int
    Sze []int
    Cnt int
}
func NewUF(n int) *UF {
    par := make([]int, n)
    sze := make([]int, n)
    for i := range par {
        par[i] = i
        sze[i] = 1
    }
    cnt := 0 
    return &UF{par, sze, cnt}
}
func (uf *UF) CntUnions() int {
    return uf.Cnt
}

func (uf *UF) Connected(v, w int) bool {
    vRoot, wRoot := uf.Find(v), uf.Find(w)
    return vRoot == wRoot
}

func (uf *UF) Find(v int) int {
    for v != uf.Par[v] {
        uf.Par[v] = uf.Par[uf.Par[v]]
        v = uf.Par[v]
    }
    return v
}

func (uf *UF) Union(v, w int) {
    rootV, rootW := uf.Find(v), uf.Find(w)
    if rootV == rootW { return }
    if uf.Sze[rootV] <= uf.Sze[rootW] {
        uf.Sze[rootV] += uf.Sze[rootW]
        uf.Par[rootW] = rootV
    } else {
        uf.Sze[rootW] += uf.Sze[rootV]
        uf.Par[rootV] = rootW
    }
    uf.Cnt++ //cnt of unions
}

func minimumCost(n int, connections [][]int) int {
    uf := NewUF(n+1) //1 <= xi, yi <= n
    sort.Sort(ByWeight(connections))
    
    res := 0
    for _, conn := range connections {
        cityA, cityB := conn[0], conn[1]
        cost := conn[2]
        if !uf.Connected(cityA,cityB) {
            uf.Union(cityA,cityB)
            res += cost
            if uf.CntUnions() == n-1 {
                return res
            }
        }
    }
    
    return -1
}

