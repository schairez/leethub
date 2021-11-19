//union find v
//time: O(n^2); finding root of node V is pseudoconstant time with path compression
//space: O(n)

func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    uf := NewUF(n)
    visited := make([][]bool, n)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    for cityAId:=0; cityAId < n; cityAId++ {
        for cityBId:=0; cityBId < n; cityBId++ {
            if !visited[cityAId][cityBId] && isConnected[cityAId][cityBId] == 1 {
                uf.Union(cityAId, cityBId)
            }
        }
    }
    return uf.cnt
} 
    
type UnionFind struct {
    par []int
    sze []int
    cnt int
}

//O(n) time
func NewUF(numNodes int) *UnionFind {
    par := make([]int, numNodes)
    sze := make([]int, numNodes)
    for i := range par {
        par[i] = i
        sze[i] = 1
    }
    uf := &UnionFind{par, sze, numNodes}
    return uf
}

//O(1) time
func (uf *UnionFind) FindRoot(u int) int {
    for uf.par[u] != u {
        uf.par[u] = uf.par[uf.par[u]]
        u = uf.par[u]
    }
    return u
}


func (uf *UnionFind) Union(u, v int) {
    rootU, rootV := uf.FindRoot(u), uf.FindRoot(v)
    if rootU == rootV {
        return
    }
    if uf.sze[rootU] <= uf.sze[rootV] { 
        uf.sze[rootV] += uf.sze[rootU]
        uf.par[rootU] = rootV
    } else {//uf.sze[rootU] > uf.sze[rootV]
        uf.sze[rootU] += uf.sze[rootV]
        uf.par[rootV] = rootU
    }
    uf.cnt--
}
    
    
    
    
    
//dfs version    
//n*n matrcityAx
//time: O(n*n) ~ O(n^2)
//space: O(n)
func findCircleNumDFS(isConnected [][]int) int {
    cityVisited := make([]int, len(isConnected))
    providenceCnt := 0
    m := len(cityVisited)
    n := len(isConnected[0])
    var dfs func(cityA int)
    dfs = func(cityA int) {
        for cityB := 0; cityB < n; cityB++ {
            if isConnected[cityA][cityB] == 1 && cityVisited[cityB] == 0 {
                cityVisited[cityB] = 1
                dfs(cityB)
                
            }
        }
    }
    for cityA :=0; cityA < m; cityA++ { 
        if cityVisited[cityA] == 0 {
            dfs(cityA)
            providenceCnt++
        }
  
    }    
    
    return providenceCnt
    
}