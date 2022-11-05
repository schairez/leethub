
func earliestAcq(logs [][]int, n int) int {
    // sort the logs first
    // init uf w/ n var
    uf := newUF(n)
    // slice sort
    sort.Slice(logs, func(i, j int) bool {
        return logs[i][0] <= logs[j][0]
    })
    friendsLeft := n-1
    for _, log := range logs {
        // logs[i] = [timestampi, xi, yi]
        timestampi, xi, yi := log[0], log[1], log[2]
        if uf.Union(xi, yi) {
            friendsLeft -= 1
        }
        if friendsLeft == 0 {
            return timestampi
        }
    }
    return -1
}


type UF struct {
    par []int
    sze []int
}

func newUF(n int) *UF {
    par := make([]int, n)
    sze := make([]int, n)
    for i := 0; i < n; i++ {
        par[i] = i
        sze[i] = 1
    }
    return &UF{par, sze}
}


func (uf *UF) FindRoot(node int) int {
    for node != uf.par[node] {
        node = uf.par[uf.par[node]]
        node = uf.par[node]
    }
    return node
}

func (uf *UF) Union(nodeA, nodeB int) bool {
    // union based on sze val
    rootA, rootB := uf.FindRoot(nodeA), uf.FindRoot(nodeB)
    rootASze, rootBSze := uf.sze[rootA], uf.sze[rootB]
    if rootA == rootB {
        return false
    }
    if rootASze >= rootBSze {
        uf.par[rootB] = rootA 
        uf.sze[nodeA] += rootBSze
    } else {
        uf.par[rootA] = rootB
        uf.sze[rootB] += rootASze
    }
    return true
} 



