//union find via weighted quick find approach


func isBipartite(graph [][]int) bool {
    n := len(graph)
    if n < 2 { 
        if len(graph[0]) == 0 { return true }
        return true 
    }
    //if len(graph[0]) == 0 { return true }
    uf := NewUF(n)
    for node := range graph {
        if len(graph[node]) == 0 { continue }
        parentNodeA := uf.Root(node)
        parentNodeB := uf.Root(graph[node][0])
        if parentNodeA == parentNodeB { return false }
        for _, neighbor := range graph[node][1:] {
            parentNodeC := uf.Root(neighbor)
            if parentNodeA == parentNodeC { return false } 
            uf.Union(parentNodeB, parentNodeC)
        }
    }
    return true
}


type UF struct {
    cnt int //count of disj sets
    id []int //parent
    sze []int //rank by size
}


//O(N) creation runtime
func NewUF(n int) *UF {
    uf := &UF{cnt: n, id: make([]int, n), sze: make([]int, n) }
    for i := range uf.id {
        uf.id[i] = i
    }
    for i := range uf.sze {
        uf.sze[i] = 1
    }
    return uf
}


//get root; path compression O(logN)
func (uf *UF) Root(p int) int {
    for uf.id[p] != p {
        uf.id[p] = uf.id[uf.id[p]]
        p = uf.id[p]
    }
    return p
}
func (uf *UF) Union(p, q int) {
    rootP := uf.Root(p)
    rootQ := uf.Root(q)
    if rootP == rootQ { return }
    if uf.sze[rootP] < uf.sze[rootQ] {
        uf.id[rootP] = rootQ
        uf.sze[rootQ] += uf.sze[rootP]
    } else {
        uf.id[rootQ] = rootP
        uf.sze[rootP] += uf.sze[rootQ]
    }
    uf.cnt--
}
//return number of individual sets (connected components)
func (uf *UF) Count() int {
    return uf.cnt
}




